package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User Info"
// @Success 201 {object} User
// @Failure 400 {object} fiber.Map
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
	return c.SendString("Create User")
}

// GetCurrentUser godoc
// @Summary Get current user
// @Description Retrieve information about the currently authenticated user
// @Tags users
// @Produce json
// @Success 200 {object} User
// @Failure 401 {object} fiber.Map
// @Router /users [get]
func GetCurrentUser(c *fiber.Ctx) error {
	return c.SendString("Get Current User")
}

// DeleteCurrentUser godoc
// @Summary Delete current user
// @Description Delete the currently authenticated user
// @Tags users
// @Produce json
// @Success 200 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Router /users [delete]
func DeleteCurrentUser(c *fiber.Ctx) error {
	return c.SendString("Delete Current User")
}