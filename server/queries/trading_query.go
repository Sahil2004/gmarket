package queries

import (
	"database/sql"
	"time"

	"github.com/Sahil2004/gmarket/server/models"
	"github.com/google/uuid"
)

type TradingQueries struct {
	*sql.DB
}

func (db *TradingQueries) EnsureTradingAccount(userID string) (models.TradingAccount, error) {
	account := models.TradingAccount{}
	query := `INSERT INTO trading_accounts (user_id) VALUES ($1)
		ON CONFLICT (user_id) DO NOTHING`
	_, err := db.Exec(query, userID)
	if err != nil {
		return account, err
	}
	return db.GetTradingAccount(userID)
}

func (db *TradingQueries) GetTradingAccount(userID string) (models.TradingAccount, error) {
	account := models.TradingAccount{}
	query := `SELECT user_id, cash_balance, margin_used, updated_at FROM trading_accounts WHERE user_id = $1`
	err := db.QueryRow(query, userID).Scan(&account.UserID, &account.CashBalance, &account.MarginUsed, &account.UpdatedAt)
	return account, err
}

func (db *TradingQueries) UpdateTradingAccount(userID string, cashBalance, marginUsed float64) error {
	query := `UPDATE trading_accounts SET cash_balance = $1, margin_used = $2, updated_at = NOW() WHERE user_id = $3`
	_, err := db.Exec(query, cashBalance, marginUsed, userID)
	return err
}

func (db *TradingQueries) ListBankAccounts(userID string) ([]models.BankAccount, error) {
	query := `SELECT id, user_id, bank_name, account_number, ifsc, nickname, created_at
		FROM bank_accounts WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := make([]models.BankAccount, 0)
	for rows.Next() {
		var a models.BankAccount
		if err := rows.Scan(&a.ID, &a.UserID, &a.BankName, &a.AccountNumber, &a.IFSC, &a.Nickname, &a.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}

func (db *TradingQueries) CreateBankAccount(userID, bankName, accountNumber, ifsc, nickname string) (models.BankAccount, error) {
	account := models.BankAccount{}
	var nick *string
	if nickname != "" {
		nick = &nickname
	}
	query := `INSERT INTO bank_accounts (user_id, bank_name, account_number, ifsc, nickname)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, bank_name, account_number, ifsc, nickname, created_at`
	err := db.QueryRow(query, userID, bankName, accountNumber, ifsc, nick).
		Scan(&account.ID, &account.UserID, &account.BankName, &account.AccountNumber, &account.IFSC, &account.Nickname, &account.CreatedAt)
	return account, err
}

func (db *TradingQueries) DeleteBankAccount(userID, bankID string) error {
	query := `DELETE FROM bank_accounts WHERE id = $1 AND user_id = $2`
	result, err := db.Exec(query, bankID, userID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (db *TradingQueries) AddFundTransaction(userID string, bankID *uuid.UUID, txType string, amount float64) error {
	query := `INSERT INTO fund_transactions (user_id, bank_account_id, type, amount) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, userID, bankID, txType, amount)
	return err
}

func (db *TradingQueries) ListHoldingsByProduct(userID, productType string) ([]models.Holding, error) {
	query := `SELECT id, user_id, symbol, exchange, product_type, quantity, avg_price, updated_at
		FROM holdings WHERE user_id = $1 AND product_type = $2 AND quantity > 0 ORDER BY symbol`
	rows, err := db.Query(query, userID, productType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	holdings := make([]models.Holding, 0)
	for rows.Next() {
		var h models.Holding
		if err := rows.Scan(&h.ID, &h.UserID, &h.Symbol, &h.Exchange, &h.ProductType, &h.Quantity, &h.AvgPrice, &h.UpdatedAt); err != nil {
			return nil, err
		}
		holdings = append(holdings, h)
	}
	return holdings, nil
}

func (db *TradingQueries) GetHolding(userID, symbol, exchange, productType string) (models.Holding, error) {
	h := models.Holding{}
	query := `SELECT id, user_id, symbol, exchange, product_type, quantity, avg_price, updated_at
		FROM holdings WHERE user_id = $1 AND symbol = $2 AND exchange = $3 AND product_type = $4`
	err := db.QueryRow(query, userID, symbol, exchange, productType).
		Scan(&h.ID, &h.UserID, &h.Symbol, &h.Exchange, &h.ProductType, &h.Quantity, &h.AvgPrice, &h.UpdatedAt)
	return h, err
}

func (db *TradingQueries) UpsertHolding(userID, symbol, exchange, productType string, quantity int, avgPrice float64) error {
	query := `INSERT INTO holdings (user_id, symbol, exchange, product_type, quantity, avg_price, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		ON CONFLICT (user_id, symbol, exchange, product_type)
		DO UPDATE SET quantity = $5, avg_price = $6, updated_at = NOW()`
	_, err := db.Exec(query, userID, symbol, exchange, productType, quantity, avgPrice)
	return err
}

func (db *TradingQueries) CreateOrder(order models.Order) (models.Order, error) {
	created := models.Order{}
	query := `INSERT INTO orders (
		user_id, symbol, exchange, side, product_type, order_type, quantity, price, stop_loss,
		status, filled_quantity, margin_required
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,'open',0,$10)
	RETURNING id, user_id, symbol, exchange, side, product_type, order_type, quantity, price, stop_loss,
		status, filled_quantity, margin_required, created_at, executed_at`
	err := db.QueryRow(query,
		order.UserID, order.Symbol, order.Exchange, order.Side, order.ProductType, order.OrderType,
		order.Quantity, order.Price, order.StopLoss, order.MarginRequired,
	).Scan(
		&created.ID, &created.UserID, &created.Symbol, &created.Exchange, &created.Side, &created.ProductType,
		&created.OrderType, &created.Quantity, &created.Price, &created.StopLoss, &created.Status,
		&created.FilledQuantity, &created.MarginRequired, &created.CreatedAt, &created.ExecutedAt,
	)
	return created, err
}

func (db *TradingQueries) ListOrdersByStatus(userID, status string) ([]models.Order, error) {
	query := `SELECT id, user_id, symbol, exchange, side, product_type, order_type, quantity, price, stop_loss,
		status, filled_quantity, margin_required, created_at, executed_at
		FROM orders WHERE user_id = $1`
	args := []interface{}{userID}
	if status != "" && status != "all" {
		query += ` AND status = $2`
		args = append(args, status)
	}
	query += ` ORDER BY created_at DESC LIMIT 100`

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]models.Order, 0)
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(
			&o.ID, &o.UserID, &o.Symbol, &o.Exchange, &o.Side, &o.ProductType, &o.OrderType,
			&o.Quantity, &o.Price, &o.StopLoss, &o.Status, &o.FilledQuantity, &o.MarginRequired,
			&o.CreatedAt, &o.ExecutedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (db *TradingQueries) GetOpenOrders(userID string) ([]models.Order, error) {
	return db.ListOrdersByStatus(userID, "open")
}

func (db *TradingQueries) GetOrder(userID, orderID string) (models.Order, error) {
	o := models.Order{}
	query := `SELECT id, user_id, symbol, exchange, side, product_type, order_type, quantity, price, stop_loss,
		status, filled_quantity, margin_required, created_at, executed_at
		FROM orders WHERE id = $1 AND user_id = $2`
	err := db.QueryRow(query, orderID, userID).Scan(
		&o.ID, &o.UserID, &o.Symbol, &o.Exchange, &o.Side, &o.ProductType, &o.OrderType,
		&o.Quantity, &o.Price, &o.StopLoss, &o.Status, &o.FilledQuantity, &o.MarginRequired,
		&o.CreatedAt, &o.ExecutedAt,
	)
	return o, err
}

func (db *TradingQueries) MarkOrderExecuted(orderID string, executedAt time.Time) error {
	query := `UPDATE orders SET status = 'executed', filled_quantity = quantity, executed_at = $1 WHERE id = $2`
	_, err := db.Exec(query, executedAt, orderID)
	return err
}

func (db *TradingQueries) CancelOrder(userID, orderID string) error {
	query := `UPDATE orders SET status = 'cancelled' WHERE id = $1 AND user_id = $2 AND status = 'open'`
	result, err := db.Exec(query, orderID, userID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
