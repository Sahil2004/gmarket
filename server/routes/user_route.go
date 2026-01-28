package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/Sahil2004/gmarket/server/services"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(a *fiber.App, imageService *services.ImageService) {
	userRouter := a.Group("/users")

	userRouter.Post("/", controllers.CreateUser)
	userRouter.Get("/", middlewares.AuthMiddleware, controllers.GetCurrentUser)
	userRouter.Delete("/", middlewares.AuthMiddleware, controllers.DeleteCurrentUser)
	userRouter.Patch("/", middlewares.AuthMiddleware, controllers.UpdateCurrentUser(imageService))
	userRouter.Post("/change-password", middlewares.AuthMiddleware, controllers.ChangePassword)
}
