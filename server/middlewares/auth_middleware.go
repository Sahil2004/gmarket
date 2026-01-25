package middlewares

import (
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	access_token := c.Get("access_token")
	refresh_token := c.Get("refresh_token")

	if access_token == "" || refresh_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorDTO{
			Code:	fiber.StatusUnauthorized,
			Message: "Access not allowed. Not logged in.",
			DevMessage: "Unauthorized: Missing tokens.",
		})
	}
	
	return c.Next()
}