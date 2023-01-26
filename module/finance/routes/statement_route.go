package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/config"
	"github.com/riyan-eng/api-finance/module/finance/controllers"
	"github.com/riyan-eng/api-finance/module/finance/repositories"
	"github.com/riyan-eng/api-finance/module/finance/services"
)

func StatementRoute(route fiber.Router) {
	statementRepository := repositories.NewStatementRepository(config.DB)
	statementService := services.NewStatementService(statementRepository)
	statementController := controllers.NewStatementController(statementService)

	statement := route.Group("/statement")
	statement.Get("/income", statementController.IncomeStatement)
	statement.Get("/capital", statementController.CapitalStatement)
	statement.Get("/balance_sheet", statementController.BalanceSheet)
	statement.Get("/cash_flow", statementController.CashFlow)
}
