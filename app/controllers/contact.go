package controllers

import (
	"github.com/Artem09076/go_api.git/app/models"
	"github.com/gofiber/fiber/v2"
)

func PostContact(c *fiber.Ctx) error {
	return c.JSON(models.Contact{})
}

func GetContact(c *fiber.Ctx) error {
	return c.JSON([]models.Contact{})
}

func PutContact(c *fiber.Ctx) error {
	return c.JSON(models.Contact{})

}

func DeleteContact(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
