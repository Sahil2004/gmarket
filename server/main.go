package main

import (
	_ "github.com/Sahil2004/gmarket/server/docs"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title GMarket API
// @version 1.0
// @description Backend API for GMarket
// @host localhost:3000
// @BasePath /api
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is the gmarket api.")
	})

	app.Get("/swagger", func (c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Listen(":3000")
}