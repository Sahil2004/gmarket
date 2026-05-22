package models

import (
	"time"

	"github.com/google/uuid"
)

type TradingAccount struct {
	UserID       uuid.UUID `db:"user_id"`
	CashBalance  float64   `db:"cash_balance"`
	MarginUsed   float64   `db:"margin_used"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type BankAccount struct {
	ID            uuid.UUID `db:"id"`
	UserID        uuid.UUID `db:"user_id"`
	BankName      string    `db:"bank_name"`
	AccountNumber string    `db:"account_number"`
	IFSC          string    `db:"ifsc"`
	Nickname      *string   `db:"nickname"`
	CreatedAt     time.Time `db:"created_at"`
}

type FundTransaction struct {
	ID            uuid.UUID  `db:"id"`
	UserID        uuid.UUID  `db:"user_id"`
	BankAccountID *uuid.UUID `db:"bank_account_id"`
	Type          string     `db:"type"`
	Amount        float64    `db:"amount"`
	CreatedAt     time.Time  `db:"created_at"`
}

type Holding struct {
	ID          uuid.UUID `db:"id"`
	UserID      uuid.UUID `db:"user_id"`
	Symbol      string    `db:"symbol"`
	Exchange    string    `db:"exchange"`
	ProductType string    `db:"product_type"`
	Quantity    int       `db:"quantity"`
	AvgPrice    float64   `db:"avg_price"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Order struct {
	ID              uuid.UUID  `db:"id"`
	UserID          uuid.UUID  `db:"user_id"`
	Symbol          string     `db:"symbol"`
	Exchange        string     `db:"exchange"`
	Side            string     `db:"side"`
	ProductType     string     `db:"product_type"`
	OrderType       string     `db:"order_type"`
	Quantity        int        `db:"quantity"`
	Price           float64    `db:"price"`
	StopLoss        *float64   `db:"stop_loss"`
	Status          string     `db:"status"`
	FilledQuantity  int        `db:"filled_quantity"`
	MarginRequired  float64    `db:"margin_required"`
	CreatedAt       time.Time  `db:"created_at"`
	ExecutedAt      *time.Time `db:"executed_at"`
}
