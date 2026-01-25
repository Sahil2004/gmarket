package queries

import (
	"database/sql"

	"github.com/Sahil2004/gmarket/server/models"
	"github.com/google/uuid"
)

type UserQueries struct {
	*sql.DB
}

func (db *UserQueries) GetUser(userID uuid.UUID) (models.User, error) {
	user := models.User{}
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE id = $1`
	err := db.QueryRow(query, userID).Scan(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *UserQueries) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}
	query := `SELECT id, email, name, password_hash, salt, profile_picture_url, phone_number, created_at, updated_at FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name, &user.PasswordHash, &user.Salt, &user.ProfilePictureUrl, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *UserQueries) CreateUser(user models.User) error {
	query := `INSERT INTO users (id, email, name, password_hash, salt, profile_picture_url, phone_number, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := db.Exec(query, user.ID, user.Email, user.Name, user.PasswordHash, user.Salt, user.ProfilePictureUrl, user.PhoneNumber, user.CreatedAt, user.UpdatedAt)
	return err
}
