package main

import (
	"github.com/Sahil2004/gmarket/server/routes"
	"github.com/gofiber/fiber/v2"
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

	routes.SwaggerRoute(app)

	app.Listen(":3000")
}