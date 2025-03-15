package controllers

import (
	"github.com/Artem09076/go_api.git/app/models"
	"github.com/gofiber/fiber/v2"
)

func PostGroup(c *fiber.Ctx) error {
	return c.JSON(models.Group{})
}

func GetGroup(c *fiber.Ctx) error {
	return c.JSON([]models.Group{})
}

func PutGroup(c *fiber.Ctx) error {
	return c.JSON(models.Group{})

}

func DeleteGroup(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
