package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkymfz/zcommerce-gofiber-api/database"
	"github.com/rizkymfz/zcommerce-gofiber-api/database/migrations"
	"github.com/rizkymfz/zcommerce-gofiber-api/routes"
)

func main() {
	app := fiber.New()

	//database init
	database.DatabaseInit()

	//migrate init
	migrations.Migration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello",
		})
	})

	//route init
	routes.RouteInit(app)
	app.Listen(":9000")
}
