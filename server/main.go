package main

import (
	"github.com/Sahil2004/gmarket/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title GMarket API
// @version 1.0
// @description Backend API for GMarket
// @host localhost:3000
// @securityDefinitions.apikey CookieAuth
// @in cookie
// @BasePath /api
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is the gmarket api.")
	})

	api := fiber.New()
	app.Mount("/api", api)

	api.Use(cors.New(cors.Config{
    	// Credentials require a specific origin, NOT "*"
		AllowOrigins:     "http://localhost:3000, http://localhost:8080", 
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Set-Cookie",
	}))

	routes.SwaggerRoute(api)
	routes.SessionRoute(api)
	routes.UserRoute(api)

	routes.NotFoundRoute(api)

	app.Listen(":3000")
}