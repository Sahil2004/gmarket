package controllers

import (
	"encoding/json"

	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/client"
)

// GetSymbols godoc
// @Summary Get Market Symbols
// @Description GetSymbols serves the list of market symbols from a static JSON file.
// @Tags market
// @Produce json
// @Success 200 {file} file "../data/stocks.json"
// @Failure 500 {object} dtos.ErrorDTO
// @Router /market/symbols [get]
func GetSymbols(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendFile("data/stocks.json")
}

// GetChartData godoc
// @Summary Get Chart Data for a Symbol
// @Description GetChartData fetches chart data for a given symbol from an external API.
// @Tags market
// @Accept json
// @Produce json
// @Param exchange query string true "Exchange (e.g., NSE)" default(NSE)
// @Param symbol query string true "Stock Symbol (e.g., RELIANCE)" default(RELIANCE)
// @Param range query string true "Chart Range (e.g., 1d, 5d, 1mo)" default(1d)
// @Param interval query string true "Data Interval (e.g., 1m, 5m, 1h)" default(1m)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dtos.ErrorDTO
// @Failure 500 {object} dtos.ErrorDTO
// @Router /market/chart [get]
func GetChartData(c *fiber.Ctx) error {
	exchange := c.Query("exchange")
	symbol := c.Query("symbol")
	chartRange := c.Query("range")
	interval := c.Query("interval")

	if exchange == "" || symbol == "" || chartRange == "" || interval == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusBadRequest,
			Message: "Missing required query parameters",
		})
	}

	var exchangeSymbol string
	switch exchange {
	case "NSE":
		exchangeSymbol = symbol + ".NS"
	default:
		exchangeSymbol = symbol
	}
	url := "https://query1.finance.yahoo.com/v8/finance/chart/" + exchangeSymbol + "?range=" + chartRange + "&interval=" + interval

	res, err := client.Get(url)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to fetch chart data",
			DevMessage: err.Error(),
		})
	}

	var result map[string]interface{}

	if err := json.Unmarshal(res.Body(), &result); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to parse chart data",
			DevMessage: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
