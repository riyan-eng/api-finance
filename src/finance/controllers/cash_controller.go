package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/src/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/src/finance/controllers/validator"
	"github.com/riyan-eng/api-finance/src/finance/services"
)

type CashController interface {
	CashReceipt(c *fiber.Ctx) error
	// CashPayment(c *fiber.Ctx) error
}

type cashController struct {
	CashService services.CashService
}

func NewCashController(cashService services.CashService) CashController {
	return &cashController{
		CashService: cashService,
	}
}

func (service cashController) CashReceipt(c *fiber.Ctx) error {
	cashReceiptBody := new(dto.CashReceiptReq)

	// parsing body json
	if err := c.BodyParser(cashReceiptBody); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if err := validator.CashReceipt(*cashReceiptBody); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// comunicate with service
	err := service.CashService.CashReceipt(*cashReceiptBody)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}
