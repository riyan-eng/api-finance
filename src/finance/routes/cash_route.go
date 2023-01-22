package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/database"
	"github.com/riyan-eng/api-finance/src/finance/controllers"
	"github.com/riyan-eng/api-finance/src/finance/repositories"
	"github.com/riyan-eng/api-finance/src/finance/services"
)

type RouterStruct struct {
	fiber.Router
}

func NewCashRoute(route RouterStruct) RouterStruct {
	return route
}

func (r *RouterStruct) CashRoutes(router fiber.Router) {
	cashRepository := repositories.NewCashRepository(database.DB)
	cashService := services.NewCashService(cashRepository)
	cashController := controllers.NewCashController(cashService)

	cash := r.Group("/cash")
	cash.Post("/receipt", cashController.CashReceipt)
}
