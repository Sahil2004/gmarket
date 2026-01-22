package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID					uuid.UUID   `db:"id" json:"id" validate:"required,uuid4"`
	Email				string 		`db:"email" json:"email" validate:"required,email"`
	Name  				string 		`db:"name" json:"name" validate:"required,min=2,max=100"`
	PasswordHash 		string 		`db:"password_hash" json:"-"`
	Salt				string 		`db:"salt" json:"-"`
	ProfilePictureUrl	 string 	 `db:"profile_picture_url" json:"profile_picture_url"`
	PhoneNumber			string		`db:"phone_number" json:"phone_number"`
	CreatedAt			time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt			time.Time	`db:"updated_at" json:"updated_at"`
}