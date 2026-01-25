package queries

import (
	"database/sql"

	"github.com/Sahil2004/gmarket/server/models"
)

type SessionQueries struct {
	*sql.DB
}

func (db *SessionQueries) CreateSession(session models.Session) error {
	query := `INSERT INTO sessions (user_id, refresh_token, created_at, expires_at) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, session.UserID, session.RefreshToken, session.CreatedAt, session.ExpiresAt)
	return err
}
