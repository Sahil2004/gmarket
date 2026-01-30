package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func MarketRoute(a *fiber.App) {
	marketRouter := a.Group("/market")

	marketRouter.Get("/symbols", middlewares.AuthMiddleware, controllers.GetSymbols)
	marketRouter.Post("/symbols/status", middlewares.AuthMiddleware, controllers.GetSymbolStatus)
	marketRouter.Get("/chart", middlewares.AuthMiddleware, controllers.GetChartData)
	marketRouter.Get("/depth", middlewares.AuthMiddleware, controllers.GetMarketDepth)
}
