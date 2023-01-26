package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/module/finance/services"
)

type StatementController interface {
	IncomeStatement(c *fiber.Ctx) error
	CapitalStatement(c *fiber.Ctx) error
	BalanceSheet(c *fiber.Ctx) error
	CashFlow(c *fiber.Ctx) error
}

type statementService struct {
	Service services.StatementService
}

func NewStatementController(sS services.StatementService) StatementController {
	return &statementService{
		Service: sS,
	}
}

func (service *statementService) IncomeStatement(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}

func (service *statementService) CapitalStatement(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}

func (service *statementService) BalanceSheet(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}

func (service *statementService) CashFlow(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}
