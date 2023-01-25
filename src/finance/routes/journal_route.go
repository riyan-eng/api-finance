package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/config"
	"github.com/riyan-eng/api-finance/src/finance/controllers"
	"github.com/riyan-eng/api-finance/src/finance/repositories"
	"github.com/riyan-eng/api-finance/src/finance/services"
)

func JournalRoutes(router fiber.Router) {
	journalRepository := repositories.NewJournalRepository(config.DB)
	journalService := services.NewJournalService(journalRepository)
	journalController := controllers.NewJournalController(journalService)

	journal := router.Group("/journal")
	journal.Post("/cash_receipt", journalController.CashReceipt)
	journal.Post("/cash_payment", journalController.CashPayment)
	journal.Post("/sales", journalController.Sales)
	journal.Post("/purchase", journalController.Purchase)
	journal.Post("/general", journalController.General)
}
