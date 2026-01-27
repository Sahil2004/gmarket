package controllers

import (
	"time"

	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/models"
	"github.com/Sahil2004/gmarket/server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body dtos.UserRegistrationDTO true "User Registration Data"
// @Success 201 {object} dtos.UserDTO
// @Failure 400 {object} dtos.ErrorDTO
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to connect to database",
			DevMessage: err.Error(),
		})
	}
	userData := &dtos.UserRegistrationDTO{}
	if err := c.BodyParser(userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid request body",
			DevMessage: err.Error(),
		})
	}
	passwordHash, salt := utils.HashPassword(userData.Password)
	user := &models.User{
		ID:           uuid.New(),
		Email:        userData.Email,
		Name:         userData.Name,
		PasswordHash: passwordHash,
		Salt:         salt,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := db.CreateUser(*user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to create user",
			DevMessage: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetCurrentUser godoc
// @Summary Get current user
// @Description Retrieve information about the currently authenticated user
// @Tags users
// @Produce json
// @Success 200 {object} dtos.UserDTO
// @Failure 401 {object} dtos.ErrorDTO
// @Router /users [get]
func GetCurrentUser(c *fiber.Ctx) error {
	return c.JSON(c.UserContext().Value("user"))
}

// DeleteCurrentUser godoc
// @Summary Delete current user
// @Description Delete the currently authenticated user
// @Tags users
// @Produce json
// @Success 200
// @Failure 401 {object} dtos.ErrorDTO
// @Router /users [delete]
func DeleteCurrentUser(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to connect to database",
			DevMessage: err.Error(),
		})
	}
	if err := db.DeleteUser(user.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to delete user",
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

	return c.SendStatus(fiber.StatusOK)
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change the password of the currently authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param password body dtos.ChangePasswordDTO true "Change Password Data"
// @Success 200
// @Failure 400 {object} dtos.ErrorDTO
// @Failure 401 {object} dtos.ErrorDTO
// @Failure 500 {object} dtos.ErrorDTO
// @Router /users/change-password [post]
func ChangePassword(c *fiber.Ctx) error {
	passwordData := dtos.ChangePasswordDTO{}
	if err := c.BodyParser(&passwordData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid request body",
			DevMessage: err.Error(),
		})
	}

	userID := c.UserContext().Value("user").(dtos.UserDTO).ID

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to connect to database",
			DevMessage: err.Error(),
		})
	}

	userData, err := db.GetUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to retrieve user data",
			DevMessage: err.Error(),
		})
	}

	if !utils.VerifyNewPassword(passwordData.OldPassword, passwordData.NewPassword, userData.Salt, userData.PasswordHash) {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Old password is incorrect or new password is the same as the old password",
			DevMessage: "Old password is incorrect or new password is the same as the old password",
		})
	}

	newHash, newSalt := utils.HashPassword(passwordData.NewPassword)

	updatedAt := time.Now().Format(time.RFC3339)

	if err := db.UpdateUserPassword(userID, newHash, newSalt, updatedAt); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to update password",
			DevMessage: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// UpdateCurrentUser godoc
// @Summary Update current user information
// @Description Update the information of the currently authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dtos.UpdateUserDTO true "User Update Data"
// @Success 200 {object} dtos.UserDTO
// @Failure 400 {object} dtos.ErrorDTO
// @Failure 401 {object} dtos.ErrorDTO
// @Failure 500 {object} dtos.ErrorDTO
// @Router /users [patch]
func UpdateCurrentUser(c *fiber.Ctx) error {
	userData := dtos.UpdateUserDTO{}
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid request body",
			DevMessage: err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to connect to database",
			DevMessage: err.Error(),
		})
	}

	userID := c.UserContext().Value("user").(dtos.UserDTO).ID
	updatedAt := time.Now().Format(time.RFC3339)

	if err := db.UpdateUserDetails(userID, userData.Email, userData.Name, userData.ProfilePictureUrl, userData.PhoneNumber, updatedAt); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to update user details",
			DevMessage: err.Error(),
		})
	}

	user, err := db.GetUser(userID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to retrieve updated user data",
			DevMessage: err.Error(),
		})
	}

	newUser := dtos.UserDTO{
		ID:                user.ID,
		Email:             user.Email,
		Name:              user.Name,
		ProfilePictureUrl: user.ProfilePictureUrl,
		PhoneNumber:       user.PhoneNumber,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}

	newAccessToken, err := utils.GenerateAccessToken(newUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to generate access token",
			DevMessage: err.Error(),
		})
	}

	accessCookie := new(fiber.Cookie)
	accessCookie.Name = "access_token"
	accessCookie.Value = newAccessToken
	accessCookie.Expires = time.Now().Add(15 * time.Minute)
	accessCookie.HTTPOnly = true
	accessCookie.Path = "/"
	accessCookie.SameSite = "Lax"
	accessCookie.Secure = false // ! Set to true in production with HTTPS

	c.Cookie(accessCookie)

	return c.Status(fiber.StatusOK).JSON(newUser)
}
