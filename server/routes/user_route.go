package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(a *fiber.App) {
	userRouter := a.Group("/users")

	userRouter.Post("/", controllers.CreateUser)
	userRouter.Get("/", middlewares.AuthMiddleware, controllers.GetCurrentUser)
	userRouter.Delete("/", middlewares.AuthMiddleware, controllers.DeleteCurrentUser)
}