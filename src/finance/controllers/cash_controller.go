package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/src/finance/controllers/dto"
)

type CashControllers interface {
	CashReceipt(c *fiber.Ctx) error
	CashPayment(c *fiber.Ctx) error
}

func CashReceipt(c *fiber.Ctx) error {
	cashReceiptBody := new(dto.CashReceiptReq)

	// parsing body json
	if err := c.BodyParser(cashReceiptBody); err != nil {
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
