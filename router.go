package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/riyan-eng/api-finance/module/finance/routes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	routes.JournalRoutes(api)
	routes.NeracaRoutes(api)
}
