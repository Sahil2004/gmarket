package controllers

import (
	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/utils"
	"github.com/gofiber/fiber/v2"
)

type AlgoController struct {
	Queries *database.Queries
}

func NewAlgoController(queries *database.Queries) *AlgoController {
	return &AlgoController{Queries: queries}
}

func (ac *AlgoController) GetConfig(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	symbol := c.Query("symbol")
	exchange := c.Query("exchange")
	if symbol == "" || exchange == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: "symbol and exchange required"})
	}
	cfg, err := utils.GetAlgoConfig(ac.Queries, user.ID.String(), symbol, exchange)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{Message: "Failed to load algo config"})
	}
	return c.JSON(cfg)
}

func (ac *AlgoController) SaveConfig(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	body := dtos.AlgoConfigDTO{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: "Invalid request body"})
	}
	if body.Symbol == "" || body.Exchange == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: "symbol and exchange required"})
	}
	cfg, err := utils.SaveAlgoConfig(ac.Queries, user.ID.String(), body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: err.Error()})
	}
	return c.JSON(cfg)
}

func (ac *AlgoController) GetIndicators(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	symbol := c.Query("symbol")
	exchange := c.Query("exchange")
	rangeKey := c.Query("range", "1D")
	if symbol == "" || exchange == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: "symbol and exchange required"})
	}
	indicators, err := utils.GetAlgoIndicators(ac.Queries, user.ID.String(), symbol, exchange, rangeKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: err.Error()})
	}
	return c.JSON(indicators)
}
