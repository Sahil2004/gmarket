package controllers

import "github.com/gofiber/fiber/v2"

// GetSymbols godoc
// @Summary Get Market Symbols
// @Description GetSymbols serves the list of market symbols from a static JSON file.
// @Tags Market
// @Produce json
// @Success 200 {file} file "../data/stocks.json"
// @Failure 500 {object} dtos.ErrorDTO
// @Router /market/symbols [get]
func GetSymbols(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendFile("data/stocks.json")
}
