package dtos

type CreateSessionDTO struct {
	Email   string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"strongpassword123"`
}

type SessionDTO struct {
	User	UserDTO `json:"user"`
}