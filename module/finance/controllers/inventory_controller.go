package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/module/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/module/finance/services"
	"github.com/riyan-eng/api-finance/util"
)

type InventoryContoller interface {
	Increase(c *fiber.Ctx) error
	Decrease(c *fiber.Ctx) error
	InventoryBalance(c *fiber.Ctx) error
}

type inventoryService struct {
	InventoryService services.InventoryService
}

func NewInventoryController(iS services.InventoryService) InventoryContoller {
	return &inventoryService{
		InventoryService: iS,
	}
}

func (service *inventoryService) Increase(c *fiber.Ctx) error {
	bodyJson := new(dto.IncreaseInventoryReq)

	if err := c.BodyParser(&bodyJson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	if err := util.Validate(*bodyJson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	if err := service.InventoryService.IncreaseStock(bodyJson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}

func (service *inventoryService) Decrease(c *fiber.Ctx) error {
	bodyJson := new(dto.DecreaseInventoryReq)

	if err := c.BodyParser(&bodyJson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	if err := util.Validate(*bodyJson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	if err := service.InventoryService.DecreaseStock(bodyJson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}

func (service *inventoryService) InventoryBalance(c *fiber.Ctx) error {
	return nil
}
