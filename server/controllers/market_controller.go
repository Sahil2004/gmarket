package controllers

import (
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/utils"
	"github.com/gofiber/fiber/v2"
)

type MarketController struct{}

func NewMarketController() *MarketController {
	return &MarketController{}
}

// GetSymbols godoc
// @Summary Get Market Symbols
// @Description GetSymbols serves the list of market symbols from a static JSON file.
// @Tags market
// @Produce json
// @Success 200 {file} file "../data/stocks.json"
// @Failure 500 {object} dtos.ErrorDTO
// @Router /market/symbols [get]
func (mc *MarketController) GetSymbols(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendFile("data/stocks.json")
}

// GetMarketDepth godoc
// @Summary Get Market Depth Data
// @Description GetMarketDepth fetches the latest price, bid, and ask levels for a given symbol.
// @Tags market
// @Accept json
// @Produce json
// @Param exchange query string true "Exchange (e.g., NSE)" default(NSE)
// @Param symbol query string true "Stock Symbol (e.g., RELIANCE)" default(RELIANCE)
// @Success 200 {object} dtos.MarketDepthDTO
// @Failure 400 {object} dtos.ErrorDTO
// @Failure 500 {object} dtos.ErrorDTO
// @Router /market/depth [get]
func (mc *MarketController) GetMarketDepth(c *fiber.Ctx) error {
	exchange := c.Query("exchange")
	symbol := c.Query("symbol")

	if exchange == "" || symbol == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusBadRequest,
			Message: "Missing required query parameters",
		})
	}

	ltp, _, bids, asks, err := utils.GetMarketData(symbol, exchange)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to fetch market depth data",
			DevMessage: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"symbol":   symbol,
		"exchange": exchange,
		"ltp":      ltp,
		"bids":     bids,
		"asks":     asks,
	})
}

// GetSymbolStatus godoc
// @Summary Get Symbol Statuses
// @Description GetSymbolStatus fetches the latest trading price and last close price for a list of symbols.
// @Tags market
// @Accept json
// @Produce json
// @Param symbols body dtos.SymbolListDTO true "List of Symbols"
// @Success 200 {object} dtos.SymbolStatusDTO
// @Failure 400 {object} dtos.ErrorDTO
// @Failure 500 {object} dtos.ErrorDTO
// @Router /market/symbols/status [post]
func (mc *MarketController) GetSymbolStatus(c *fiber.Ctx) error {
	symbols := dtos.SymbolListDTO{}
	if err := c.BodyParser(&symbols); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusBadRequest,
			Message:    "Invalid request body",
			DevMessage: err.Error(),
		})
	}

	statuses, err := utils.GetSymbolsStatus(symbols.Symbols)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to fetch symbol statuses",
			DevMessage: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(statuses)
}
