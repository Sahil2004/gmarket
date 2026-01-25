package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	RefreshToken	string   `db:"refresh_token" json:"refresh_token" validate:"required"`
	UserID			uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid4"`
	CreatedAt		time.Time `db:"created_at" json:"created_at" validate:"required"`
	ExpiresAt		time.Time `db:"expires_at" json:"expires_at"`
}