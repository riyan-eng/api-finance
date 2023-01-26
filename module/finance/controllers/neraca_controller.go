package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/module/finance/services"
)

type NeracaController interface {
	TrialBalance(c *fiber.Ctx) error
	TrialBalanceAfterAdjustment(c *fiber.Ctx) error
	BalanceSheet(c *fiber.Ctx) error
}

type neracaService struct {
	NeracaService services.NeracaService
}

func NewNeracaController(nC services.NeracaService) NeracaController {
	return &neracaService{
		NeracaService: nC,
	}
}

func (service neracaService) TrialBalance(c *fiber.Ctx) error {
	// trialBalanceBody:=new(dto.TrialBalanceReq)

	// parsing body json

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}

func (service neracaService) TrialBalanceAfterAdjustment(c *fiber.Ctx) error {
	return nil
}

func (service neracaService) BalanceSheet(c *fiber.Ctx) error {
	return nil
}
