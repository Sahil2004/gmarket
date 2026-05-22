package models

import (
	"time"

	"github.com/google/uuid"
)

type AlgoConfig struct {
	UserID         uuid.UUID `db:"user_id"`
	Symbol         string    `db:"symbol"`
	Exchange       string    `db:"exchange"`
	Enabled        bool      `db:"enabled"`
	RSIEnabled     bool      `db:"rsi_enabled"`
	RSIPeriod      int       `db:"rsi_period"`
	RSIOverbought  float64   `db:"rsi_overbought"`
	RSIOversold    float64   `db:"rsi_oversold"`
	MAEnabled      bool      `db:"ma_enabled"`
	MAFastPeriod   int       `db:"ma_fast_period"`
	MASlowPeriod   int       `db:"ma_slow_period"`
	UpdatedAt      time.Time `db:"updated_at"`
}
