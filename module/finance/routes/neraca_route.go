package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/config"
	"github.com/riyan-eng/api-finance/module/finance/controllers"
	"github.com/riyan-eng/api-finance/module/finance/repositories"
	"github.com/riyan-eng/api-finance/module/finance/services"
)

func NeracaRoutes(router fiber.Router) {
	neracaRepository := repositories.NewNeracaRepository(config.DB)
	neracaService := services.NewNeracaService(neracaRepository)
	neracaController := controllers.NewNeracaController(neracaService)

	neraca := router.Group("/neraca")
	neraca.Get("/coa", neracaController.ChartOfAccount)
	neraca.Get("/trial_balance", neracaController.TrialBalance)
	neraca.Get("/trial_balance_adjustment", neracaController.TrialBalanceAfterAdjustment)
}
