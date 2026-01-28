package controllers

import (
	"strconv"

	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/utils"
	"github.com/gofiber/fiber/v2"
)

// GetWatchlist godoc
// @Summary Get a user's watchlist
// @Description Retrieve the specified watchlist for the authenticated user
// @Tags watchlists
// @Produce json
// @Param watchlist_idx path int true "Watchlist Index (1-10)"
// @Success 200 {object} dtos.WatchlistDTO
// @Failure 400 {object} dtos.ErrorDTO
// @Failure 401 {object} dtos.ErrorDTO
// @Failure 500 {object} dtos.ErrorDTO
// @Router /watchlists/{watchlist_idx} [get]
func GetWatchlist(c *fiber.Ctx) error {
	idx, err := strconv.Atoi(c.Params("watchlist_idx"))

	if err != nil || idx < 1 || idx > 10 {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid watchlist index",
			DevMessage: err.Error(),
		})
	}

	user := c.UserContext().Value("user").(dtos.UserDTO)

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Database connection error",
			DevMessage: err.Error(),
		})
	}

	watchlist, err := db.GetWatchlist(user.ID.String(), idx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to retrieve watchlist",
			DevMessage: err.Error(),
		})
	}

	symbolInfos := make([]dtos.SymbolInfo, 0)

	for _, symbol := range watchlist.Symbols {
		actualSymbol, exchange := utils.ParseDBStockSymbol(symbol)
		fullSymbol := utils.GetStockSymbolWithExchange(actualSymbol, exchange)
		exchangeData, err := utils.GetExchangeData(fullSymbol, "1d", "1m")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
				Code:       fiber.StatusInternalServerError,
				Message:    "Failed to fetch symbol data",
				DevMessage: err.Error(),
			})
		}
		symbolInfos = append(symbolInfos, dtos.SymbolInfo{
			Symbol:         actualSymbol,
			Exchange:       exchange,
			LTP:            exchangeData["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["meta"].(map[string]interface{})["regularMarketPrice"].(float64),
			LastClosePrice: exchangeData["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["meta"].(map[string]interface{})["chartPreviousClose"].(float64),
		})
	}

	watchlistWithData := dtos.WatchlistDTO{
		Index:   idx,
		Symbols: symbolInfos,
	}

	return c.Status(fiber.StatusOK).JSON(watchlistWithData)
}

// AddSymbolToWatchlist godoc
// @Summary Add a symbol to a user's watchlist
// @Description Add a stock symbol to the specified watchlist for the authenticated user
// @Tags watchlists
// @Accept json
// @Produce json
// @Param watchlist_idx path int true "Watchlist Index (1-10)"
// @Param symbol body dtos.SymbolDTO true "Symbol to add"
// @Success 200 {object} dtos.SuccessDTO
// @Failure 400 {object} dtos.ErrorDTO
// @Failure 401 {object} dtos.ErrorDTO
// @Failure 500 {object} dtos.ErrorDTO
// @Router /watchlists/{watchlist_idx}/symbols [post]
func AddSymbolToWatchlist(c *fiber.Ctx) error {
	idx, err := strconv.Atoi(c.Params("watchlist_idx"))

	if err != nil || idx < 1 || idx > 10 {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid watchlist index",
			DevMessage: err.Error(),
		})
	}

	symbol := dtos.SymbolDTO{}

	if err := c.BodyParser(&symbol); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid request body",
			DevMessage: err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Database connection error",
			DevMessage: err.Error(),
		})
	}

	user := c.UserContext().Value("user").(dtos.UserDTO)

	dbSymbol := utils.GenerateDBStockSymbol(symbol.Symbol, symbol.Exchange)

	db.AddSymbolToWatchlist(user.ID.String(), idx, dbSymbol)

	return c.Status(fiber.StatusOK).JSON(dtos.SuccessDTO{
		Code:    fiber.StatusOK,
		Message: "Symbol added successfully",
	})
}

// RemoveSymbolFromWatchlist godoc
// @Summary Remove a symbol from a user's watchlist
// @Description Remove a stock symbol from the specified watchlist for the authenticated user
// @Tags watchlists
// @Accept json
// @Produce json
// @Param watchlist_idx path int true "Watchlist Index (1-10)"
// @Param symbol body dtos.SymbolDTO true "Symbol to remove"
// @Success 200 {object} dtos.SuccessDTO
// @Failure 400 {object} dtos.ErrorDTO
// @Failure 401 {object} dtos.ErrorDTO
// @Failure 500 {object} dtos.ErrorDTO
// @Router /watchlists/{watchlist_idx}/symbols [delete]
func RemoveSymbolFromWatchlist(c *fiber.Ctx) error {
	idx, err := strconv.Atoi(c.Params("watchlist_idx"))

	if err != nil || idx < 1 || idx > 10 {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid watchlist index",
			DevMessage: err.Error(),
		})
	}

	symbol := dtos.SymbolDTO{}

	if err := c.BodyParser(&symbol); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid request body",
			DevMessage: err.Error(),
		})
	}

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Database connection error",
			DevMessage: err.Error(),
		})
	}

	user := c.UserContext().Value("user").(dtos.UserDTO)
	dbSymbol := utils.GenerateDBStockSymbol(symbol.Symbol, symbol.Exchange)

	db.RemoveSymbolFromWatchlist(user.ID.String(), idx, dbSymbol)

	return c.Status(fiber.StatusOK).JSON(dtos.SuccessDTO{
		Code:    fiber.StatusOK,
		Message: "Symbol removed successfully",
	})
}
