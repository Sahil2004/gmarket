package controllers

import (
	"time"

	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jaevor/go-nanoid"
	"golang.org/x/crypto/argon2"
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
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to connect to the database",
		})
	}
	userData := &dtos.UserRegistrationDTO{}
	if err := c.BodyParser(userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}
	salt, _ := nanoid.Standard(8)
	user := &models.User{
		ID: uuid.New(),
		Email: userData.Email,
		Name: userData.Name,
		PasswordHash: string(argon2.Key([]byte(userData.Password), []byte(salt()), 3, 32*1024, 4, 32)),
		Salt: salt(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.CreateUser(*user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetCurrentUser godoc
// @Summary Get current user
// @Description Retrieve information about the currently authenticated user
// @Tags users
// @Produce json
// @Failure 401 {object} dtos.ErrorDTO
// @Router /users [get]
func GetCurrentUser(c *fiber.Ctx) error {
	return c.SendString("Get Current User")
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
	return c.SendString("Delete Current User")
}