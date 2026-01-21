package dtos

type ErrorDTO struct {
	Code	int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}