package models

type Watchlist struct {
	UserID       string   `db:"user_id" json:"user_id"`
	WatchlistIdx int      `db:"watchlist_idx" json:"watchlist_idx"`
	Symbols      []string `db:"symbols" json:"symbols"`
	UpdatedAt    string   `db:"updated_at" json:"updated_at"`
}
