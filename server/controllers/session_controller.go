package controllers

import "github.com/gofiber/fiber/v2"

func CreateSession(c *fiber.Ctx) error {
	return c.SendString("Create Session")
}

func DeleteCurrentSession(c *fiber.Ctx) error {
	return c.SendString("Delete Current Session")
}