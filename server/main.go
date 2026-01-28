package main

import (
	"log"

	"github.com/Sahil2004/gmarket/server/routes"
	"github.com/Sahil2004/gmarket/server/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	api := fiber.New()
	app.Mount("/api", api)

	api.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:4200",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Set-Cookie",
	}))

	imageService, err := services.NewImageService()
	if err != nil {
		log.Fatal(err)
	}

	routes.SwaggerRoute(api)
	routes.SessionRoute(api)
	routes.UserRoute(api, imageService)
	routes.MarketRoute(api)
	routes.WatchlistRoute(api)

	routes.NotFoundRoute(api)

	app.Listen(":3000")
}
