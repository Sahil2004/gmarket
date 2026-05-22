package queries

import (
	"database/sql"

	"github.com/Sahil2004/gmarket/server/models"
)

type AlgoQueries struct {
	*sql.DB
}

func (db *AlgoQueries) GetAlgoConfig(userID, symbol, exchange string) (models.AlgoConfig, error) {
	cfg := models.AlgoConfig{}
	query := `SELECT user_id, symbol, exchange, enabled, rsi_enabled, rsi_period, rsi_overbought, rsi_oversold,
		ma_enabled, ma_fast_period, ma_slow_period, updated_at
		FROM algo_configs WHERE user_id = $1 AND symbol = $2 AND exchange = $3`
	err := db.QueryRow(query, userID, symbol, exchange).Scan(
		&cfg.UserID, &cfg.Symbol, &cfg.Exchange, &cfg.Enabled, &cfg.RSIEnabled, &cfg.RSIPeriod,
		&cfg.RSIOverbought, &cfg.RSIOversold, &cfg.MAEnabled, &cfg.MAFastPeriod, &cfg.MASlowPeriod, &cfg.UpdatedAt,
	)
	return cfg, err
}

func (db *AlgoQueries) UpsertAlgoConfig(cfg models.AlgoConfig) (models.AlgoConfig, error) {
	query := `INSERT INTO algo_configs (
		user_id, symbol, exchange, enabled, rsi_enabled, rsi_period, rsi_overbought, rsi_oversold,
		ma_enabled, ma_fast_period, ma_slow_period, updated_at
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,NOW())
	ON CONFLICT (user_id, symbol, exchange) DO UPDATE SET
		enabled = $4, rsi_enabled = $5, rsi_period = $6, rsi_overbought = $7, rsi_oversold = $8,
		ma_enabled = $9, ma_fast_period = $10, ma_slow_period = $11, updated_at = NOW()
	RETURNING user_id, symbol, exchange, enabled, rsi_enabled, rsi_period, rsi_overbought, rsi_oversold,
		ma_enabled, ma_fast_period, ma_slow_period, updated_at`
	out := models.AlgoConfig{}
	err := db.QueryRow(query,
		cfg.UserID, cfg.Symbol, cfg.Exchange, cfg.Enabled, cfg.RSIEnabled, cfg.RSIPeriod,
		cfg.RSIOverbought, cfg.RSIOversold, cfg.MAEnabled, cfg.MAFastPeriod, cfg.MASlowPeriod,
	).Scan(
		&out.UserID, &out.Symbol, &out.Exchange, &out.Enabled, &out.RSIEnabled, &out.RSIPeriod,
		&out.RSIOverbought, &out.RSIOversold, &out.MAEnabled, &out.MAFastPeriod, &out.MASlowPeriod, &out.UpdatedAt,
	)
	return out, err
}
