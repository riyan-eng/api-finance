package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/config"
	"github.com/riyan-eng/api-finance/module/finance/controllers"
	"github.com/riyan-eng/api-finance/module/finance/repositories"
	"github.com/riyan-eng/api-finance/module/finance/services"
)

func InventoryRoutes(router fiber.Router) {
	inventoryRepository := repositories.NewInventoryRepository(config.DB)
	inventoryService := services.NewInventoryService(inventoryRepository)
	inventoryController := controllers.NewInventoryController(inventoryService)

	inventory := router.Group("/inventory")
	inventory.Post("/increase", inventoryController.Increase)
	inventory.Post("/decrease", inventoryController.Decrease)
	inventory.Post("/balance", inventoryController.InventoryBalance)
}
