package dtos

type CreateSessionDTO struct {
	Email   string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"strongpassword123"`
}

type SessionDTO struct {
	ID		string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	User	UserDTO `json:"user"`
}