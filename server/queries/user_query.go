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
	query := `SELECT id, email, password_hash, salt, created_at, updated_at FROM users WHERE id = $1`
	err := db.QueryRow(query, userID).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Salt, &user.CreatedAt, &user.UpdatedAt)
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

func (db *UserQueries) DeleteUser(userId uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := db.Exec(query, userId)
	return err
}

func (db *UserQueries) UpdateUserPassword(userId uuid.UUID, newPasswordHash string, newSalt string, updatedAt string) error {
	query := `UPDATE users SET password_hash = $1, salt = $2, updated_at = $3 WHERE id = $4;`
	_, err := db.Exec(query, newPasswordHash, newSalt, updatedAt, userId)
	return err
}
