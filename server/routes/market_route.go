package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func MarketRoute(a *fiber.App) {
	marketRouter := a.Group("/market")

	marketRouter.Get("/symbols", middlewares.AuthMiddleware, controllers.GetSymbols)
	marketRouter.Get("/chart", middlewares.AuthMiddleware, controllers.GetChartData)
}
