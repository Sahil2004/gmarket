package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func WatchlistRoute(a *fiber.App) {
	watchlistRouter := a.Group("/watchlists")

	watchlistRouter.Get("/:watchlist_idx", middlewares.AuthMiddleware, controllers.GetWatchlist)
}
