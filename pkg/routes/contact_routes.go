package routes

import (
	"github.com/Artem09076/go_api.git/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func ContactRoutes(a *fiber.App) {
	route := a.Group("/api/v1/contact")

	route.Get("/", controllers.GetContact)
	route.Post("/", controllers.PostContact)
	route.Put("/", controllers.PutContact)
	route.Delete("/", controllers.DeleteContact)
}
