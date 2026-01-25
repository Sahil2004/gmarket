package queries

import (
	"database/sql"

	"github.com/Sahil2004/gmarket/server/models"
)

type WatchlistQueries struct {
	*sql.DB
}

func (db *WatchlistQueries) GetWatchlist(userID string, watchlistIdx int) (models.Watchlist, error) {
	watchlist := models.Watchlist{}
	query := `SELECT user_id, watchlist_idx, symbols, updated_at FROM watchlists WHERE user_id = $1 AND watchlist_idx = $2`
	err := db.QueryRow(query, userID, watchlistIdx).Scan(&watchlist.UserID, &watchlist.WatchlistIdx, &watchlist.Symbols, &watchlist.UpdatedAt)
	if err != nil {
		return watchlist, err
	}
	return watchlist, nil
}

func (db *WatchlistQueries) AddSymbolToWatchlist(userID string, watchlistIdx int, symbol string) error {
	query := `UPDATE watchlists SET symbols = array_append(symbols, $1), updated_at = NOW() WHERE user_id = $2 AND watchlist_idx = $3 AND NOT ($1 = ANY(symbols))`
	_, err := db.Exec(query, symbol, userID, watchlistIdx)
	return err
}

func (db *WatchlistQueries) RemoveSymbolFromWatchlist(userID string, watchlistIdx int, symbol string) error {
	query := `UPDATE watchlists SET symbols = array_remove(symbols, $1), updated_at = NOW() WHERE user_id = $2 AND watchlist_idx = $3`
	_, err := db.Exec(query, symbol, userID, watchlistIdx)
	return err
}
