package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SessionRoute(a *fiber.App, queries *database.Queries) {
	sessionController := controllers.NewSessionController(queries)
	sessionRouter := a.Group("/sessions")

	sessionRouter.Post("/", sessionController.CreateSession)
	sessionRouter.Delete("/", middlewares.AuthMiddleware, sessionController.DeleteCurrentSession)
}
