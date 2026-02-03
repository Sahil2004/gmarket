package controllers

import (
	"time"

	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/models"
	"github.com/Sahil2004/gmarket/server/utils"
	"github.com/gofiber/fiber/v2"
)

type SessionController struct {
	Queries *database.Queries
}

func NewSessionController(queries *database.Queries) *SessionController {
	return &SessionController{
		Queries: queries,
	}
}

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
func (sc *SessionController) CreateSession(c *fiber.Ctx) error {
	loginData := &dtos.CreateSessionDTO{}
	if err := c.BodyParser(loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid login details",
			DevMessage: err.Error(),
		})
	}

	userData, err := sc.Queries.GetUserByEmail(loginData.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Unable to get that user",
			DevMessage: err.Error(),
		})
	}

	if !utils.ValidatePassword(loginData.Password, userData.Salt, userData.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusUnauthorized,
			Message:    "Invalid credentials",
			DevMessage: "Password and the hash do not match.",
		})
	}

	user := dtos.UserDTO{
		ID:                userData.ID,
		Email:             userData.Email,
		Name:              userData.Name,
		ProfilePictureUrl: userData.ProfilePictureUrl,
		PhoneNumber:       userData.PhoneNumber,
		CreatedAt:         userData.CreatedAt,
		UpdatedAt:         userData.UpdatedAt,
	}

	accessTokenCookie := new(fiber.Cookie)
	accessTokenCookie.Name = "access_token"
	accessTokenCookie.Value, err = utils.GenerateAccessToken(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to generate access token",
			DevMessage: err.Error(),
		})
	}

	accessTokenCookie.MaxAge = 15 * 60 // 15 minutes
	accessTokenCookie.HTTPOnly = true
	accessTokenCookie.Path = "/"
	accessTokenCookie.SameSite = "Lax"
	accessTokenCookie.Secure = false // ! set true in production with HTTPS

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to generate refresh token",
			DevMessage: err.Error(),
		})
	}

	refreshTokenCookie := new(fiber.Cookie)
	refreshTokenCookie.Name = "refresh_token"
	refreshTokenCookie.Value = refreshToken
	refreshTokenCookie.MaxAge = 7 * 24 * 60 * 60 // 7 days
	refreshTokenCookie.HTTPOnly = true
	refreshTokenCookie.Path = "/"
	refreshTokenCookie.SameSite = "Lax"
	refreshTokenCookie.Secure = false // ! set true in production with HTTPS

	session := &models.Session{
		RefreshToken: refreshToken,
		UserID:       user.ID,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	if err := sc.Queries.CreateSession(*session); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to create session",
			DevMessage: err.Error(),
		})
	}

	c.Cookie(accessTokenCookie)
	c.Cookie(refreshTokenCookie)

	return c.Status(fiber.StatusCreated).JSON(dtos.SessionDTO{
		User: user,
	})
}

// DeleteCurrentSession godoc
// @Summary Delete the current session (logout)
// @Description Delete the current session for the authenticated user
// @Tags sessions
// @Produce json
// @Success 200 {object} dtos.SuccessDTO
// @Failure 401 {object} dtos.ErrorDTO
// @Router /sessions [delete]
func (sc *SessionController) DeleteCurrentSession(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")

	if err := sc.Queries.DeleteSession(refreshToken); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to delete session",
			DevMessage: err.Error(),
		})
	}

	expiredAccess := new(fiber.Cookie)
	expiredAccess.Name = "access_token"
	expiredAccess.Value = ""
	expiredAccess.MaxAge = -1
	expiredAccess.Path = "/"
	expiredAccess.HTTPOnly = true
	expiredAccess.SameSite = "Lax"
	expiredAccess.Secure = false

	expiredRefresh := new(fiber.Cookie)
	expiredRefresh.Name = "refresh_token"
	expiredRefresh.Value = ""
	expiredRefresh.MaxAge = -1
	expiredRefresh.Path = "/"
	expiredRefresh.HTTPOnly = true
	expiredRefresh.SameSite = "Lax"
	expiredRefresh.Secure = false

	c.Cookie(expiredAccess)
	c.Cookie(expiredRefresh)

	return c.Status(fiber.StatusOK).JSON(dtos.SuccessDTO{
		Code:    fiber.StatusOK,
		Message: "Logged out successfully",
	})
}
