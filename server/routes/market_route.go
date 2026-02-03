package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func MarketRoute(a *fiber.App) {
	marketController := controllers.NewMarketController()
	marketRouter := a.Group("/market")

	marketRouter.Get("/symbols", middlewares.AuthMiddleware, marketController.GetSymbols)
	marketRouter.Post("/symbols/status", middlewares.AuthMiddleware, marketController.GetSymbolStatus)
	marketRouter.Get("/depth", middlewares.AuthMiddleware, marketController.GetMarketDepth)
}
