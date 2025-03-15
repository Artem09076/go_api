package routes

import (
	"github.com/Artem09076/go_api.git/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func GroupRoutes(a *fiber.App) {
	routes := a.Group("/api/v1/group")

	routes.Get("/", controllers.GetGroup)
	routes.Post("/", controllers.PostGroup)
	routes.Put("/", controllers.PutGroup)
	routes.Delete("/", controllers.DeleteGroup)

}
