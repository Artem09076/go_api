package main

import (
	"log"

	"github.com/Artem09076/go_api.git/pkg/routes"
	"github.com/Artem09076/go_api.git/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("AAAAAAAAA")
	})
	database.PostgreSQLConnection()
	routes.ContactRoutes(app)
	routes.GroupRoutes(app)

	if err := app.Listen(":6080"); err != nil {
		log.Fatalln(err)
	}
}
