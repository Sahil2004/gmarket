package middlewares

import (
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code: fiber.StatusBadRequest,
			Message: "Missing or malformed JWT",
			DevMessage: err.Error(),
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorDTO{
		Code: fiber.StatusUnauthorized,
		Message: "Invalid or expired JWT",
		DevMessage: err.Error(),
	})
}

func JwtMiddleware() func(*fiber.Ctx) error {
	config := jwtware.Config{
		ContextKey: "user",
		ErrorHandler: jwtErrorHandler,
	}
	return jwtware.New(config)
}