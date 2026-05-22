CREATE TABLE algo_configs (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    symbol VARCHAR(50) NOT NULL,
    exchange VARCHAR(10) NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT false,
    rsi_enabled BOOLEAN NOT NULL DEFAULT true,
    rsi_period INT NOT NULL DEFAULT 14 CHECK (rsi_period BETWEEN 2 AND 100),
    rsi_overbought DECIMAL(5, 2) NOT NULL DEFAULT 70 CHECK (rsi_overbought BETWEEN 50 AND 100),
    rsi_oversold DECIMAL(5, 2) NOT NULL DEFAULT 30 CHECK (rsi_oversold BETWEEN 0 AND 50),
    ma_enabled BOOLEAN NOT NULL DEFAULT true,
    ma_fast_period INT NOT NULL DEFAULT 9 CHECK (ma_fast_period BETWEEN 2 AND 200),
    ma_slow_period INT NOT NULL DEFAULT 21 CHECK (ma_slow_period BETWEEN 2 AND 200),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, symbol, exchange),
    CHECK (ma_fast_period < ma_slow_period)
);
