package utils

import (
	"database/sql"
	"errors"
	"math"
	"time"

	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/models"
)

const intradayMarginRate = 0.2

func RoundMoney(value float64) float64 {
	return math.Round(value*100) / 100
}

func CalcMarginRequired(price float64, qty int, productType string) float64 {
	total := price * float64(qty)
	if productType == "intraday" {
		return RoundMoney(total * intradayMarginRate)
	}
	return RoundMoney(total)
}

func CalcAvailable(cashBalance, marginUsed float64) float64 {
	return RoundMoney(cashBalance - marginUsed)
}

func HoldingToDTO(h models.Holding, ltp, lastClose float64) dtos.HoldingDTO {
	currentValue := RoundMoney(ltp * float64(h.Quantity))
	totalPnL := RoundMoney((ltp - h.AvgPrice) * float64(h.Quantity))
	dayPnL := RoundMoney((ltp - lastClose) * float64(h.Quantity))
	return dtos.HoldingDTO{
		Symbol:       h.Symbol,
		Exchange:     h.Exchange,
		ProductType:  h.ProductType,
		Quantity:     h.Quantity,
		AvgPrice:     h.AvgPrice,
		LTP:          ltp,
		LastClose:    lastClose,
		PnL:          totalPnL,
		DayPnL:       dayPnL,
		TotalPnL:     totalPnL,
		CurrentValue: currentValue,
	}
}

func OrderToDTO(o models.Order) dtos.OrderDTO {
	return dtos.OrderDTO{
		ID:             o.ID.String(),
		Symbol:         o.Symbol,
		Exchange:       o.Exchange,
		Side:           o.Side,
		ProductType:    o.ProductType,
		OrderType:      o.OrderType,
		Quantity:       o.Quantity,
		Price:          o.Price,
		StopLoss:       o.StopLoss,
		Status:         o.Status,
		FilledQuantity: o.FilledQuantity,
		MarginRequired: o.MarginRequired,
		CreatedAt:      o.CreatedAt,
		ExecutedAt:     o.ExecutedAt,
	}
}

func AccountToDTO(a models.TradingAccount) dtos.TradingAccountDTO {
	return dtos.TradingAccountDTO{
		CashBalance: a.CashBalance,
		MarginUsed:  a.MarginUsed,
		Available:   CalcAvailable(a.CashBalance, a.MarginUsed),
		UpdatedAt:   a.UpdatedAt,
	}
}

func CanExecuteOrder(o models.Order, ltp float64) bool {
	if o.OrderType == "market" {
		return true
	}
	if o.Side == "buy" {
		return ltp <= o.Price
	}
	return ltp >= o.Price
}

func ProcessOpenOrders(queries *database.Queries, userID string) error {
	orders, err := queries.GetOpenOrders(userID)
	if err != nil {
		return err
	}
	for _, order := range orders {
		ltp, _, _, _, err := GetMarketData(order.Symbol, order.Exchange)
		if err != nil {
			continue
		}
		if !CanExecuteOrder(order, ltp) {
			continue
		}
		if err := ExecuteOrder(queries, userID, order, ltp); err != nil {
			continue
		}
	}
	return nil
}

func ExecuteOrder(queries *database.Queries, userID string, order models.Order, executionPrice float64) error {
	account, err := queries.EnsureTradingAccount(userID)
	if err != nil {
		return err
	}

	execPrice := RoundMoney(executionPrice)
	if order.OrderType == "limit" {
		execPrice = order.Price
	}

	if order.Side == "buy" {
		if err := executeBuy(queries, userID, order, account, execPrice); err != nil {
			return err
		}
	} else {
		if err := executeSell(queries, userID, order, account, execPrice); err != nil {
			return err
		}
	}

	return queries.MarkOrderExecuted(order.ID.String(), time.Now())
}

func executeBuy(queries *database.Queries, userID string, order models.Order, account models.TradingAccount, execPrice float64) error {
	margin := CalcMarginRequired(execPrice, order.Quantity, order.ProductType)
	if order.ProductType == "regular" {
		if account.CashBalance < margin {
			return errors.New("insufficient funds")
		}
		account.CashBalance = RoundMoney(account.CashBalance - margin)
	} else {
		if CalcAvailable(account.CashBalance, account.MarginUsed) < margin {
			return errors.New("insufficient funds")
		}
		account.MarginUsed = RoundMoney(account.MarginUsed + margin)
	}

	if err := queries.UpdateTradingAccount(userID, account.CashBalance, account.MarginUsed); err != nil {
		return err
	}

	h, err := queries.GetHolding(userID, order.Symbol, order.Exchange, order.ProductType)
	qty := order.Quantity
	avg := execPrice
	if err == nil {
		totalQty := h.Quantity + order.Quantity
		avg = RoundMoney((h.AvgPrice*float64(h.Quantity) + execPrice*float64(order.Quantity)) / float64(totalQty))
		qty = totalQty
	} else if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return queries.UpsertHolding(userID, order.Symbol, order.Exchange, order.ProductType, qty, avg)
}

func executeSell(queries *database.Queries, userID string, order models.Order, account models.TradingAccount, execPrice float64) error {
	h, err := queries.GetHolding(userID, order.Symbol, order.Exchange, order.ProductType)
	if err != nil {
		return errors.New("insufficient holdings")
	}
	if h.Quantity < order.Quantity {
		return errors.New("insufficient holdings")
	}

	proceeds := RoundMoney(execPrice * float64(order.Quantity))
	account.CashBalance = RoundMoney(account.CashBalance + proceeds)

	marginRelease := CalcMarginRequired(h.AvgPrice, order.Quantity, order.ProductType)
	if order.ProductType == "intraday" {
		account.MarginUsed = RoundMoney(math.Max(0, account.MarginUsed-marginRelease))
	}

	if err := queries.UpdateTradingAccount(userID, account.CashBalance, account.MarginUsed); err != nil {
		return err
	}

	newQty := h.Quantity - order.Quantity
	if newQty <= 0 {
		return queries.UpsertHolding(userID, order.Symbol, order.Exchange, order.ProductType, 0, 0)
	}
	return queries.UpsertHolding(userID, order.Symbol, order.Exchange, order.ProductType, newQty, h.AvgPrice)
}

func PlaceOrder(queries *database.Queries, userID string, req dtos.PlaceOrderDTO) (dtos.OrderDTO, error) {
	if req.Quantity <= 0 {
		return dtos.OrderDTO{}, errors.New("invalid quantity")
	}
	if req.ProductType != "regular" && req.ProductType != "intraday" {
		return dtos.OrderDTO{}, errors.New("invalid product type")
	}
	if req.Side != "buy" && req.Side != "sell" {
		return dtos.OrderDTO{}, errors.New("invalid side")
	}
	if req.OrderType != "limit" && req.OrderType != "market" {
		return dtos.OrderDTO{}, errors.New("invalid order type")
	}

	ltp, _, _, _, err := GetMarketData(req.Symbol, req.Exchange)
	if err != nil {
		return dtos.OrderDTO{}, err
	}

	price := req.Price
	if req.OrderType == "market" {
		price = ltp
	}
	if price <= 0 {
		return dtos.OrderDTO{}, errors.New("invalid price")
	}

	account, err := queries.EnsureTradingAccount(userID)
	if err != nil {
		return dtos.OrderDTO{}, err
	}

	margin := CalcMarginRequired(price, req.Quantity, req.ProductType)
	if req.Side == "buy" && CalcAvailable(account.CashBalance, account.MarginUsed) < margin {
		return dtos.OrderDTO{}, errors.New("insufficient funds")
	}
	if req.Side == "sell" {
		h, err := queries.GetHolding(userID, req.Symbol, req.Exchange, req.ProductType)
		if err != nil || h.Quantity < req.Quantity {
			return dtos.OrderDTO{}, errors.New("insufficient holdings")
		}
	}

	userUUID := account.UserID
	order := models.Order{
		UserID:         userUUID,
		Symbol:         req.Symbol,
		Exchange:       req.Exchange,
		Side:           req.Side,
		ProductType:    req.ProductType,
		OrderType:      req.OrderType,
		Quantity:       req.Quantity,
		Price:          RoundMoney(price),
		StopLoss:       req.StopLoss,
		MarginRequired: margin,
	}

	created, err := queries.CreateOrder(order)
	if err != nil {
		return dtos.OrderDTO{}, err
	}

	if CanExecuteOrder(created, ltp) {
		if err := ExecuteOrder(queries, userID, created, ltp); err == nil {
			created, _ = queries.GetOrder(userID, created.ID.String())
		}
	}

	return OrderToDTO(created), nil
}

func BuildTradingSnapshot(queries *database.Queries, userID string) (dtos.TradingSnapshotDTO, error) {
	_ = ProcessOpenOrders(queries, userID)

	account, err := queries.EnsureTradingAccount(userID)
	if err != nil {
		return dtos.TradingSnapshotDTO{}, err
	}

	holdings, _ := queries.ListHoldingsByProduct(userID, "regular")
	positions, _ := queries.ListHoldingsByProduct(userID, "intraday")
	openOrders, _ := queries.ListOrdersByStatus(userID, "open")
	executedOrders, _ := queries.ListOrdersByStatus(userID, "executed")
	banks, _ := queries.ListBankAccounts(userID)

	var totalHoldingsPnL float64
	holdingDTOs := make([]dtos.HoldingDTO, 0, len(holdings))
	for _, h := range holdings {
		ltp, lcp, _, _, _ := GetMarketData(h.Symbol, h.Exchange)
		dto := HoldingToDTO(h, ltp, lcp)
		totalHoldingsPnL += dto.TotalPnL
		holdingDTOs = append(holdingDTOs, dto)
	}

	var totalPositionsDayPnL float64
	positionDTOs := make([]dtos.HoldingDTO, 0, len(positions))
	for _, h := range positions {
		ltp, lcp, _, _, _ := GetMarketData(h.Symbol, h.Exchange)
		dto := HoldingToDTO(h, ltp, lcp)
		totalPositionsDayPnL += dto.DayPnL
		positionDTOs = append(positionDTOs, dto)
	}

	openDTOs := make([]dtos.OrderDTO, 0, len(openOrders))
	for _, o := range openOrders {
		openDTOs = append(openDTOs, OrderToDTO(o))
	}

	executedDTOs := make([]dtos.OrderDTO, 0, len(executedOrders))
	for _, o := range executedOrders {
		executedDTOs = append(executedDTOs, OrderToDTO(o))
	}

	bankDTOs := make([]dtos.BankAccountDTO, 0, len(banks))
	for _, b := range banks {
		nick := ""
		if b.Nickname != nil {
			nick = *b.Nickname
		}
		bankDTOs = append(bankDTOs, dtos.BankAccountDTO{
			ID:            b.ID.String(),
			BankName:      b.BankName,
			AccountNumber: b.AccountNumber,
			IFSC:          b.IFSC,
			Nickname:      nick,
			CreatedAt:     b.CreatedAt,
		})
	}

	return dtos.TradingSnapshotDTO{
		Account:              AccountToDTO(account),
		Holdings:             holdingDTOs,
		Positions:            positionDTOs,
		TotalHoldingsPnL:     RoundMoney(totalHoldingsPnL),
		TotalPositionsDayPnL: RoundMoney(totalPositionsDayPnL),
		OpenOrders:           openDTOs,
		ExecutedOrders:       executedDTOs,
		BankAccounts:         bankDTOs,
	}, nil
}

func OrderPreview(queries *database.Queries, userID string, req dtos.PlaceOrderDTO) (dtos.OrderPreviewDTO, error) {
	ltp, _, _, _, err := GetMarketData(req.Symbol, req.Exchange)
	if err != nil {
		return dtos.OrderPreviewDTO{}, err
	}
	price := req.Price
	if req.OrderType == "market" {
		price = ltp
	}
	account, err := queries.EnsureTradingAccount(userID)
	if err != nil {
		return dtos.OrderPreviewDTO{}, err
	}
	margin := CalcMarginRequired(price, req.Quantity, req.ProductType)
	return dtos.OrderPreviewDTO{
		Symbol:         req.Symbol,
		Exchange:       req.Exchange,
		Side:           req.Side,
		ProductType:    req.ProductType,
		Quantity:       req.Quantity,
		Price:          RoundMoney(price),
		LTP:            ltp,
		MarginRequired: margin,
		Available:      CalcAvailable(account.CashBalance, account.MarginUsed),
		OrderValue:     RoundMoney(price * float64(req.Quantity)),
	}, nil
}
