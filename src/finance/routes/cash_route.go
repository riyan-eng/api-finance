package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/src/finance/controllers"
)

type RouterStruct struct {
	fiber.Router
}

func (r *RouterStruct) CashRoutes(router fiber.Router) {
	// cashRepository:=""
	// cashService:=""
	// cashController=""

	cash := r.Group("/cash")
	cash.Post("/receipt", controllers.CashReceipt)
}
