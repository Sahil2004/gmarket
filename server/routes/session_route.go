package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SessionRoute(a *fiber.App) {
	sessionRouter := a.Group("/sessions")

	sessionRouter.Post("/", controllers.CreateSession)
	sessionRouter.Delete("/", middlewares.AuthMiddleware, controllers.DeleteCurrentSession)
}