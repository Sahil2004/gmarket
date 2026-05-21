package controllers

import (
	"database/sql"

	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TradingController struct {
	Queries *database.Queries
}

func NewTradingController(queries *database.Queries) *TradingController {
	return &TradingController{Queries: queries}
}

func (tc *TradingController) GetSnapshot(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	snapshot, err := utils.BuildTradingSnapshot(tc.Queries, user.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:       fiber.StatusInternalServerError,
			Message:    "Failed to load trading data",
			DevMessage: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(snapshot)
}

func (tc *TradingController) GetOrderPreview(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	req := dtos.PlaceOrderDTO{
		Symbol:      c.Query("symbol"),
		Exchange:    c.Query("exchange"),
		Side:        c.Query("side"),
		ProductType: c.Query("product_type", "regular"),
		OrderType:   c.Query("order_type", "limit"),
		Quantity:    c.QueryInt("quantity", 1),
		Price:       c.QueryFloat("price", 0),
	}
	preview, err := utils.OrderPreview(tc.Queries, user.ID.String(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(preview)
}

func (tc *TradingController) PlaceOrder(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	req := dtos.PlaceOrderDTO{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}
	order, err := utils.PlaceOrder(tc.Queries, user.ID.String(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}

func (tc *TradingController) CancelOrder(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	orderID := c.Params("order_id")
	if err := tc.Queries.CancelOrder(user.ID.String(), orderID); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(dtos.ErrorDTO{
				Code:    fiber.StatusNotFound,
				Message: "Order not found or not cancellable",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to cancel order",
		})
	}
	return c.Status(fiber.StatusOK).JSON(dtos.SuccessDTO{Message: "Order cancelled"})
}

func (tc *TradingController) ListBankAccounts(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	_, _ = tc.Queries.EnsureTradingAccount(user.ID.String())
	snapshot, err := utils.BuildTradingSnapshot(tc.Queries, user.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{Message: "Failed to load bank accounts"})
	}
	return c.Status(fiber.StatusOK).JSON(snapshot.BankAccounts)
}

func (tc *TradingController) CreateBankAccount(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	req := dtos.CreateBankAccountDTO{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: "Invalid request body"})
	}
	account, err := tc.Queries.CreateBankAccount(
		user.ID.String(), req.BankName, req.AccountNumber, req.IFSC, req.Nickname,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{Message: "Failed to add bank account"})
	}
	nick := ""
	if account.Nickname != nil {
		nick = *account.Nickname
	}
	return c.Status(fiber.StatusCreated).JSON(dtos.BankAccountDTO{
		ID: account.ID.String(), BankName: account.BankName,
		AccountNumber: account.AccountNumber, IFSC: account.IFSC,
		Nickname: nick, CreatedAt: account.CreatedAt,
	})
}

func (tc *TradingController) DeleteBankAccount(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	if err := tc.Queries.DeleteBankAccount(user.ID.String(), c.Params("bank_id")); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dtos.ErrorDTO{Message: "Bank account not found"})
	}
	return c.Status(fiber.StatusOK).JSON(dtos.SuccessDTO{Message: "Bank account removed"})
}

func (tc *TradingController) DepositFunds(c *fiber.Ctx) error {
	return tc.transferFunds(c, "deposit")
}

func (tc *TradingController) WithdrawFunds(c *fiber.Ctx) error {
	return tc.transferFunds(c, "withdraw")
}

func (tc *TradingController) transferFunds(c *fiber.Ctx, txType string) error {
	user := c.UserContext().Value("user").(dtos.UserDTO)
	req := dtos.FundTransferDTO{}
	if err := c.BodyParser(&req); err != nil || req.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: "Invalid amount"})
	}

	account, err := tc.Queries.EnsureTradingAccount(user.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{Message: "Failed to load account"})
	}

	var bankID *uuid.UUID
	if req.BankAccountID != "" {
		id, err := uuid.Parse(req.BankAccountID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: "Invalid bank account"})
		}
		bankID = &id
	}

	if txType == "withdraw" {
		if account.CashBalance < req.Amount {
			return c.Status(fiber.StatusBadRequest).JSON(dtos.ErrorDTO{Message: "Insufficient balance"})
		}
		account.CashBalance = utils.RoundMoney(account.CashBalance - req.Amount)
	} else {
		account.CashBalance = utils.RoundMoney(account.CashBalance + req.Amount)
	}

	if err := tc.Queries.UpdateTradingAccount(user.ID.String(), account.CashBalance, account.MarginUsed); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{Message: "Failed to update balance"})
	}
	if err := tc.Queries.AddFundTransaction(user.ID.String(), bankID, txType, req.Amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.ErrorDTO{Message: "Failed to record transaction"})
	}

	return c.Status(fiber.StatusOK).JSON(utils.AccountToDTO(account))
}
