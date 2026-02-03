package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func WatchlistRoute(a *fiber.App, queries *database.Queries) {
	watchlistController := controllers.NewWatchlistController(queries)
	watchlistRouter := a.Group("/watchlists")

	watchlistRouter.Get("/:watchlist_idx", middlewares.AuthMiddleware, watchlistController.GetWatchlist)
	watchlistRouter.Post("/:watchlist_idx/symbols", middlewares.AuthMiddleware, watchlistController.AddSymbolToWatchlist)
	watchlistRouter.Delete("/:watchlist_idx/symbols", middlewares.AuthMiddleware, watchlistController.RemoveSymbolFromWatchlist)
}
