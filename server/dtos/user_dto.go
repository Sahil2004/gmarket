package dtos

import (
	"time"

	"github.com/google/uuid"
)

type UserRegistrationDTO struct {
	Email    string `json:"email" example:"john@example.com"`
	Name     string `json:"name" example:"John Doe"`
	Password string `json:"password" example:"strongpassword123"`
}

type UserDTO struct {
	ID                uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email             string    `json:"email" example:"john@example.com"`
	Name              string    `json:"name" example:"John Doe"`
	ProfilePictureUrl string    `json:"profile_picture_url" example:"https://example.com/profile.jpg"`
	PhoneNumber       string    `json:"phone_number" example:"+1234567890"`
	CreatedAt         time.Time `json:"created_at" example:"2023-10-01T12:00:00Z"`
	UpdatedAt         time.Time `json:"updated_at" example:"2023-10-01T12:00:00Z"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"old_password" example:"oldpassword123"`
	NewPassword string `json:"new_password" example:"newstrongpassword456"`
}
