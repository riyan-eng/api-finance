package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/src/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/src/finance/controllers/validator"
	"github.com/riyan-eng/api-finance/src/finance/services"
)

type CashController interface {
	CashReceipt(c *fiber.Ctx) error
	CashPayment(c *fiber.Ctx) error
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
	// fmt.Println(cashReceiptBody)

	// parsing body json
	if err := c.BodyParser(&cashReceiptBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	// fmt.Println(cashReceiptBody)

	// validate body json
	if err := validator.CashReceipt(*cashReceiptBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}
	// fmt.Println(cashReceiptBody)

	// comunicate with service
	err := service.CashService.CashReceipt(*cashReceiptBody)
	fmt.Println("--- controller ---")
	fmt.Println(err)
	fmt.Println("--- controller ---")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted cash receipt",
		"message": "ok",
	})
}

func (service cashController) CashPayment(c *fiber.Ctx) error {
	cashPaymentBody := new(dto.CashPaymentReq)

	// parsing body json
	if err := c.BodyParser(&cashPaymentBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate  body json
	if err := validator.CashPayment(*cashPaymentBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// communicate with service

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted cash payment",
		"message": "ok",
	})
}
