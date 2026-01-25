package queries

import (
	"database/sql"

	"github.com/Sahil2004/gmarket/server/models"
	"github.com/lib/pq"
)

type WatchlistQueries struct {
	*sql.DB
}

func (db *WatchlistQueries) GetWatchlist(userID string, watchlistIdx int) (models.Watchlist, error) {
	watchlist := models.Watchlist{}
	query := `SELECT user_id, watchlist_idx, symbols, updated_at FROM watchlists WHERE user_id = $1 AND watchlist_idx = $2`
	err := db.QueryRow(query, userID, watchlistIdx).Scan(&watchlist.UserID, &watchlist.WatchlistIdx, pq.Array(&watchlist.Symbols), &watchlist.UpdatedAt)
	if err != nil {
		return watchlist, err
	}
	return watchlist, nil
}

func (db *WatchlistQueries) AddSymbolToWatchlist(userID string, watchlistIdx int, symbol string) error {
	query := `INSERT INTO watchlists (user_id, watchlist_idx, symbols, updated_at) VALUES ($1, $2, ARRAY[$3], NOW()) ON CONFLICT (user_id, watchlist_idx) DO UPDATE SET symbols = array_append(watchlists.symbols, $3), updated_at = NOW() WHERE NOT ($3 = ANY(watchlists.symbols))`
	_, err := db.Exec(query, userID, watchlistIdx, symbol)
	return err
}

func (db *WatchlistQueries) RemoveSymbolFromWatchlist(userID string, watchlistIdx int, symbol string) error {
	query := `UPDATE watchlists SET symbols = array_remove(symbols, $1), updated_at = NOW() WHERE user_id = $2 AND watchlist_idx = $3`
	_, err := db.Exec(query, symbol, userID, watchlistIdx)
	return err
}
