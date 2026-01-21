package dtos

type UserRegistrationDTO struct {
	Email    string `json:"email" example:"john@example.com"`
	Name     string `json:"name" example:"John Doe"`
	Password string `json:"password" example:"strongpassword123"`
}

type UserDTO struct {
	ID	string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email    string `json:"email" example:"john@example.com"`
	Name     string `json:"name" example:"John Doe"`
	ProfilePictureUrl string `json:"profile_picture_url" example:"https://example.com/profile.jpg"`
	PhoneNumber string `json:"phone_number" example:"+1234567890"`
	CreatedAt string `json:"created_at" example:"2023-10-01T12:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2023-10-01T12:00:00Z"`
}