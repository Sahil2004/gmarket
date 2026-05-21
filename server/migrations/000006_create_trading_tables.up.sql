CREATE TABLE trading_accounts (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    cash_balance DECIMAL(18, 2) NOT NULL DEFAULT 100000.00,
    margin_used DECIMAL(18, 2) NOT NULL DEFAULT 0.00,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bank_accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    bank_name VARCHAR(100) NOT NULL,
    account_number VARCHAR(30) NOT NULL,
    ifsc VARCHAR(20) NOT NULL,
    nickname VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE fund_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    bank_account_id UUID REFERENCES bank_accounts(id) ON DELETE SET NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('deposit', 'withdraw')),
    amount DECIMAL(18, 2) NOT NULL CHECK (amount > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE holdings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    symbol VARCHAR(50) NOT NULL,
    exchange VARCHAR(10) NOT NULL,
    product_type VARCHAR(10) NOT NULL CHECK (product_type IN ('regular', 'intraday')),
    quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),
    avg_price DECIMAL(18, 2) NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, symbol, exchange, product_type)
);

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    symbol VARCHAR(50) NOT NULL,
    exchange VARCHAR(10) NOT NULL,
    side VARCHAR(4) NOT NULL CHECK (side IN ('buy', 'sell')),
    product_type VARCHAR(10) NOT NULL CHECK (product_type IN ('regular', 'intraday')),
    order_type VARCHAR(10) NOT NULL CHECK (order_type IN ('limit', 'market')),
    quantity INT NOT NULL CHECK (quantity > 0),
    price DECIMAL(18, 2) NOT NULL,
    stop_loss DECIMAL(18, 2),
    status VARCHAR(10) NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'executed', 'cancelled')),
    filled_quantity INT NOT NULL DEFAULT 0,
    margin_required DECIMAL(18, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    executed_at TIMESTAMPTZ
);

CREATE INDEX idx_orders_user_status ON orders(user_id, status);
CREATE INDEX idx_holdings_user_product ON holdings(user_id, product_type);
