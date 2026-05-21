package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AlgoRoute(a *fiber.App, queries *database.Queries) {
	ac := controllers.NewAlgoController(queries)
	algo := a.Group("/algo", middlewares.AuthMiddleware)
	algo.Get("/config", ac.GetConfig)
	algo.Put("/config", ac.SaveConfig)
	algo.Get("/indicators", ac.GetIndicators)
}
