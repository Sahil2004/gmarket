package routes

import (
	"github.com/Sahil2004/gmarket/server/controllers"
	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func TradingRoute(a *fiber.App, queries *database.Queries) {
	tc := controllers.NewTradingController(queries)
	trading := a.Group("/trading", middlewares.AuthMiddleware)

	trading.Get("/snapshot", tc.GetSnapshot)
	trading.Get("/orders/preview", tc.GetOrderPreview)
	trading.Post("/orders", tc.PlaceOrder)
	trading.Delete("/orders/:order_id", tc.CancelOrder)

	trading.Get("/banks", tc.ListBankAccounts)
	trading.Post("/banks", tc.CreateBankAccount)
	trading.Delete("/banks/:bank_id", tc.DeleteBankAccount)
	trading.Post("/funds/deposit", tc.DepositFunds)
	trading.Post("/funds/withdraw", tc.WithdrawFunds)
}
