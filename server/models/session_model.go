package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID				uuid.UUID `db:"id" json:"id" validate:"required,uuid4"`
	UserID			uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid4"`
	AccessToken		string   `db:"access_token" json:"access_token" validate:"required"`
	RefreshToken	string   `db:"refresh_token" json:"refresh_token" validate:"required"`
	CreatedAt		time.Time `db:"created_at" json:"created_at"`
	ExpiresAt		time.Time `db:"expires_at" json:"expires_at"`
}