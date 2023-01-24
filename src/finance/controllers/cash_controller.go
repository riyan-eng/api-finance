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
	Sales(c *fiber.Ctx) error
	Purchase(c *fiber.Ctx) error
	General(c *fiber.Ctx) error
}

type cashController struct {
	CashService services.CashService
}

func NewCashController(cashService services.CashService) CashController {
	return &cashController{
		CashService: cashService,
	}
}

func (service *cashController) CashReceipt(c *fiber.Ctx) error {
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

func (service *cashController) CashPayment(c *fiber.Ctx) error {
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
	if err := service.CashService.CashPayment(*cashPaymentBody); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted cash payment",
		"message": "ok",
	})
}

func (service *cashController) Sales(c *fiber.Ctx) error {
	salesBody := new(dto.SalesReq)

	// parse body json
	if err := c.BodyParser(&salesBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if err := validator.Sales(*salesBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// communicate with service
	if err := service.CashService.Sales(*salesBody); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted sales",
		"message": "ok",
	})
}

func (service *cashController) Purchase(c *fiber.Ctx) error {
	purchaseBody := new(dto.PurchaseReq)

	// parse body json
	if err := c.BodyParser(&purchaseBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if err := validator.Purchase(*purchaseBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// communicate with service
	if err := service.CashService.Purchase(*purchaseBody); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted purchase",
		"message": "ok",
	})
}

func (service *cashController) General(c *fiber.Ctx) error {
	generalBody := new(dto.GeneralReq)

	// parse body json
	if err := c.BodyParser(&generalBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if err := validator.General(*generalBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	// communicate with service
	if err := service.CashService.General(*generalBody); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "success inserted general",
		"message": "ok",
	})
}
