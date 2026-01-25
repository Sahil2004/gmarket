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

func (db *SessionQueries) GetSessionByRefreshToken(refreshToken string) (models.Session, error) {
	session := models.Session{}
	query := `SELECT refresh_token, user_id, created_at, expires_at FROM sessions WHERE refresh_token = $1`
	err := db.QueryRow(query, refreshToken).Scan(&session.RefreshToken, &session.UserID, &session.CreatedAt, &session.ExpiresAt)
	if err != nil {
		return session, err
	}
	return session, nil
}

func (db *SessionQueries) DeleteSession(refreshToken string) error {
	query := `DELETE FROM sessions WHERE refresh_token = $1`
	_, err := db.Exec(query, refreshToken)
	return err
}
