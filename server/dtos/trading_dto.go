package dtos

import "time"

type TradingAccountDTO struct {
	CashBalance  float64 `json:"cash_balance"`
	MarginUsed   float64 `json:"margin_used"`
	Available    float64 `json:"available"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type BankAccountDTO struct {
	ID            string  `json:"id"`
	BankName      string  `json:"bank_name"`
	AccountNumber string  `json:"account_number"`
	IFSC          string  `json:"ifsc"`
	Nickname      string  `json:"nickname,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateBankAccountDTO struct {
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
	IFSC          string `json:"ifsc"`
	Nickname      string `json:"nickname"`
}

type FundTransferDTO struct {
	BankAccountID string  `json:"bank_account_id"`
	Amount        float64 `json:"amount"`
}

type HoldingDTO struct {
	Symbol       string  `json:"symbol"`
	Exchange     string  `json:"exchange"`
	ProductType  string  `json:"product_type"`
	Quantity     int     `json:"quantity"`
	AvgPrice     float64 `json:"avg_price"`
	LTP          float64 `json:"ltp"`
	LastClose    float64 `json:"last_close"`
	PnL          float64 `json:"pnl"`
	DayPnL       float64 `json:"day_pnl"`
	TotalPnL     float64 `json:"total_pnl"`
	CurrentValue float64 `json:"current_value"`
}

type OrderDTO struct {
	ID             string     `json:"id"`
	Symbol         string     `json:"symbol"`
	Exchange       string     `json:"exchange"`
	Side           string     `json:"side"`
	ProductType    string     `json:"product_type"`
	OrderType      string     `json:"order_type"`
	Quantity       int        `json:"quantity"`
	Price          float64    `json:"price"`
	StopLoss       *float64   `json:"stop_loss,omitempty"`
	Status         string     `json:"status"`
	FilledQuantity int        `json:"filled_quantity"`
	MarginRequired float64    `json:"margin_required"`
	CreatedAt      time.Time  `json:"created_at"`
	ExecutedAt     *time.Time `json:"executed_at,omitempty"`
}

type PlaceOrderDTO struct {
	Symbol      string   `json:"symbol"`
	Exchange    string   `json:"exchange"`
	Side        string   `json:"side"`
	ProductType string   `json:"product_type"`
	OrderType   string   `json:"order_type"`
	Quantity    int      `json:"quantity"`
	Price       float64  `json:"price"`
	StopLoss    *float64 `json:"stop_loss"`
}

type OrderPreviewDTO struct {
	Symbol         string  `json:"symbol"`
	Exchange       string  `json:"exchange"`
	Side           string  `json:"side"`
	ProductType    string  `json:"product_type"`
	Quantity       int     `json:"quantity"`
	Price          float64 `json:"price"`
	LTP            float64 `json:"ltp"`
	MarginRequired float64 `json:"margin_required"`
	Available      float64 `json:"available"`
	OrderValue     float64 `json:"order_value"`
}

type TradingSnapshotDTO struct {
	Account            TradingAccountDTO `json:"account"`
	Holdings           []HoldingDTO      `json:"holdings"`
	Positions          []HoldingDTO      `json:"positions"`
	TotalHoldingsPnL   float64           `json:"total_holdings_pnl"`
	TotalPositionsDayPnL float64         `json:"total_positions_day_pnl"`
	OpenOrders         []OrderDTO        `json:"open_orders"`
	ExecutedOrders     []OrderDTO        `json:"executed_orders"`
	BankAccounts       []BankAccountDTO  `json:"bank_accounts"`
}
