package dtos

type ErrorDTO struct {
	Code	int    `json:"code" example:"400"`
	DevMessage string `json:"dev_message,omitempty" example:"Invalid request body format"`
	Message string `json:"message" example:"Bad Request"`
}