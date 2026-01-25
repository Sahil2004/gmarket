package middlewares

import (
	"context"
	"time"

	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthMiddleware(c *fiber.Ctx) error {
	access_token := c.Cookies("access_token")
	refresh_token := c.Cookies("refresh_token")

	if access_token == "" || refresh_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusUnauthorized,
			Message:    "Access not allowed. Not logged in.",
			DevMessage: "Unauthorized: Missing tokens.",
		})
	}

	accessTokenClaims, refreshTokenClaims, err := utils.ValidateTokens(access_token, refresh_token)

	if err != nil {
		if refreshTokenClaims == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorDTO{
				Code:       fiber.StatusUnauthorized,
				Message:    "Access not allowed. Invalid tokens.",
				DevMessage: err.Error(),
			})
		}
		if accessTokenClaims == nil {
			if claims, ok := refreshTokenClaims.Claims.(jwt.MapClaims); ok && refreshTokenClaims.Valid {
				db, err := database.OpenDBConnection()
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
						Code:       fiber.StatusInternalServerError,
						Message:    "Failed to connect to database",
						DevMessage: err.Error(),
					})
				}

				var userId uuid.UUID
				userId, _ = uuid.Parse(claims["user_id"].(string))
				user, err := db.GetUser(userId)

				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
						Code:       fiber.StatusInternalServerError,
						Message:    "Failed to get user data",
						DevMessage: err.Error(),
					})
				}

				userData := dtos.UserDTO{
					ID:                user.ID,
					Email:             user.Email,
					Name:              user.Name,
					ProfilePictureUrl: user.ProfilePictureUrl,
					PhoneNumber:       user.PhoneNumber,
					CreatedAt:         user.CreatedAt,
					UpdatedAt:         user.UpdatedAt,
				}

				newAccessToken, err := utils.GenerateAccessToken(userData)

				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
						Code:       fiber.StatusInternalServerError,
						Message:    "Failed to generate new access token",
						DevMessage: err.Error(),
					})
				}

				accessTokenCookie := new(fiber.Cookie)
				accessTokenCookie.Name = "access_token"
				accessTokenCookie.Value = newAccessToken
				accessTokenCookie.Expires = time.Now().Add(15 * time.Minute)
				accessTokenCookie.HTTPOnly = true
				accessTokenCookie.SameSite = "Lax"
				accessTokenCookie.Secure = false // ! set true in production with HTTPS

				c.Cookie(accessTokenCookie)

				userCtx := context.WithValue(c.UserContext(), "user", userData)
				c.SetUserContext(userCtx)

				return c.Next()
			}
		}
	}

	if claims, ok := accessTokenClaims.Claims.(jwt.MapClaims); ok && accessTokenClaims.Valid {
		var user dtos.UserDTO
		user = claims["user"].(dtos.UserDTO)
		userCtx := context.WithValue(c.UserContext(), "user", user)
		c.SetUserContext(userCtx)
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorDTO{
		Code:       fiber.StatusUnauthorized,
		Message:    "Access not allowed. Invalid access token.",
		DevMessage: "Unauthorized: Invalid access token claims.",
	})
}
