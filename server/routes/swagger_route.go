package routes

import (
	_ "github.com/Sahil2004/gmarket/server/docs"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SwaggerRoute(a *fiber.App) {
	router := a.Group("/swagger")

	router.Get("/", func (c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	router.Get("/*", fiberSwagger.WrapHandler)
}