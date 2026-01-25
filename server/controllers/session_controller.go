package controllers

import (
	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateSession godoc
// @Summary Create a new session (login)
// @Description Create a new session for a user
// @Tags sessions
// @Accept json
// @Produce json
// @Param user body dtos.CreateSessionDTO true "User Login Data"
// @Success 201 {object} dtos.SessionDTO
// @Failure 400 {object} dtos.ErrorDTO
// @Router /sessions [post]
func CreateSession(c *fiber.Ctx) error {
	loginData := &dtos.CreateSessionDTO{}
	if err := c.BodyParser(loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid login details",
			DevMessage: err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to connect to database",
			DevMessage: err.Error(),
		})
	}

	userData, err := db.GetUserByEmail(loginData.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:   fiber.StatusBadRequest,
			Message: "Unable to get that user",
			DevMessage: err.Error(),
		})
	}

	if !utils.ValidatePassword(loginData.Password, userData.Salt, userData.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusUnauthorized,
			Message: "Invalid credentials",
			DevMessage: "Password and the hash do not match.",
		})
	}

	user := dtos.UserDTO{
		ID:   userData.ID,
		Email: userData.Email,
		Name: userData.Name,
		ProfilePictureUrl: userData.ProfilePictureUrl,
		PhoneNumber: userData.PhoneNumber,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	accessTokenCookie := new(fiber.Cookie)
	accessTokenCookie.Name = "access_token"
	accessTokenCookie.Value, err = utils.GenerateAccessToken(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to generate access token",
			DevMessage: err.Error(),
		})
	}

	accessTokenCookie.MaxAge = 15 * 60 // 15 minutes
	accessTokenCookie.HTTPOnly = true
	accessTokenCookie.Path = "/"
	accessTokenCookie.SameSite = "Lax"
	accessTokenCookie.Secure = false // ! set true in production with HTTPS

	refreshTokenCookie := new(fiber.Cookie)
	refreshTokenCookie.Name = "refresh_token"
	refreshTokenCookie.Value, err = utils.GenerateRefreshToken(user)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to generate refresh token",
			DevMessage: err.Error(),
		})
	}

	refreshTokenCookie.MaxAge = 7 * 24 * 60 * 60 // 7 days
	refreshTokenCookie.HTTPOnly = true
	refreshTokenCookie.Path = "/"
	refreshTokenCookie.SameSite = "Lax"
	refreshTokenCookie.Secure = false // ! set true in production with HTTPS

	c.Cookie(accessTokenCookie)
	c.Cookie(refreshTokenCookie)

	return c.Status(fiber.StatusCreated).JSON(dtos.SessionDTO{
		ID: "1",
		User: user,
	})
}

// DeleteCurrentSession godoc
// @Summary Delete the current session (logout)
// @Description Delete the current session for the authenticated user
// @Tags sessions
// @Security BearerAuth
// @Produce json
// @Success 200
// @Failure 401 {object} dtos.ErrorDTO
// @Router /sessions [delete]
func DeleteCurrentSession(c *fiber.Ctx) error {
	return c.SendString("Delete Current Session")
}