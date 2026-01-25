-- define table
CREATE TABLE IF NOT EXISTS watchlists (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    watchlist_idx INT NOT NULL CHECK (watchlist_idx BETWEEN 1 AND 10),
    symbols TEXT[],
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, watchlist_idx)
);