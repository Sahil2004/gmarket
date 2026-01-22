package controllers

import "github.com/gofiber/fiber/v2"

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
	return c.SendString("Create Session")
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