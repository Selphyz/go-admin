package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-admin/src/database"
	"go-admin/src/routes"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
