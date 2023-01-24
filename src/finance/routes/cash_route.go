package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/database"
	"github.com/riyan-eng/api-finance/src/finance/controllers"
	"github.com/riyan-eng/api-finance/src/finance/repositories"
	"github.com/riyan-eng/api-finance/src/finance/services"
)

func CashRoutes(router fiber.Router) {
	cashRepository := repositories.NewCashRepository(database.DB)
	cashService := services.NewCashService(cashRepository)
	cashController := controllers.NewCashController(cashService)

	cash := router.Group("/cash")
	cash.Post("/receipt", cashController.CashReceipt)
	cash.Post("/payment", cashController.CashPayment)
	cash.Post("/sales", cashController.Sales)
	cash.Post("/purchase", cashController.Purchase)
	cash.Post("/general", cashController.General)
}
