package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/module/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/module/finance/services"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
	"github.com/riyan-eng/api-finance/util"
)

type NeracaController interface {
	TrialBalance(c *fiber.Ctx) error
	TrialBalanceAfterAdjustment(c *fiber.Ctx) error
}

type neracaService struct {
	NeracaService services.NeracaService
}

func NewNeracaController(nC services.NeracaService) NeracaController {
	return &neracaService{
		NeracaService: nC,
	}
}

func (service *neracaService) TrialBalance(c *fiber.Ctx) error {
	trialBalanceBody := new(dto.TrialBalanceReq)

	// parsing query
	if err := c.QueryParser(trialBalanceBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate query
	if err := util.Validate(trialBalanceBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	fmt.Println(trialBalanceBody)

	lala := entities.NeracaEntity{}
	neracas, err := service.NeracaService.TrialBalance(lala)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    neracas,
		"message": "ok",
	})
}

func (service *neracaService) TrialBalanceAfterAdjustment(c *fiber.Ctx) error {
	trialBalanceBody := new(dto.TrialBalanceReq)

	// parsing query
	if err := c.QueryParser(trialBalanceBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate query
	if err := util.Validate(trialBalanceBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "fail",
		})
	}

	fmt.Println(trialBalanceBody)

	lala := entities.NeracaEntity{}
	neracas, err := service.NeracaService.TrialBalanceAfterAdjustment(lala)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    neracas,
		"message": "ok",
	})
}
