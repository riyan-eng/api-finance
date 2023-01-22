package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/riyan-eng/api-finance/src/finance/routes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	routes.CashRoutes(api)
}
