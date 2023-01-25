package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/riyan-eng/api-finance/config"
)

func init() {
	config.LoadEnv()
	config.ConnectDb()
}

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("help me!!!")
		return err
	})

	SetupRoutes(app)

	app.Listen(":3000")
}
