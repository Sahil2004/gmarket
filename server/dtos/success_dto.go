package dtos

type SuccessDTO struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"OK"`
}
