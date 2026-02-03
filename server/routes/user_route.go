package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/Sahil2004/gmarket/server/services"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(a *fiber.App, queries *database.Queries, imageService *services.ImageService) {
	userController := controllers.NewUserController(queries, imageService)
	userRouter := a.Group("/users")

	userRouter.Post("/", userController.CreateUser)
	userRouter.Get("/", middlewares.AuthMiddleware, userController.GetCurrentUser)
	userRouter.Delete("/", middlewares.AuthMiddleware, userController.DeleteCurrentUser)
	userRouter.Patch("/", middlewares.AuthMiddleware, userController.UpdateCurrentUser)
	userRouter.Post("/change-password", middlewares.AuthMiddleware, userController.ChangePassword)
	userRouter.Get("/is-authenticated", middlewares.AuthMiddleware, userController.GetAuthStatus)
}
