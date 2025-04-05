package controllers

import (
	"github.com/Artem09076/go_api.git/app/models"
	"github.com/Artem09076/go_api.git/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostContact(c *fiber.Ctx) error {
	db := database.DB
	contact := new(models.Contact)
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Invalid credentials", "data": err})
	}

	contact.ID = uuid.New()
	if err := db.Table("contact").Create(&contact).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create contact", "data": err})
	}
	if contact.GroupID != nil && *contact.GroupID == uuid.Nil {
		contact.GroupID = nil
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Created contact", "data": contact})
}

func GetContact(c *fiber.Ctx) error {
	db := database.DB
	var contacts []models.Contact
	if err := db.Table("contact").Preload("Phones").Find(&contacts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(contacts)
}

func PutContact(c *fiber.Ctx) error {
	var contact models.Contact
	id := c.Params("id")
	db := database.DB
	db.Table("contact").Find(&contact, "id = ?", id)
	if contact.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"err": "Contact not found"})
	}
	var updateContactData map[string]any

	if err := c.BodyParser(&updateContactData); err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Invalid credentials", "data": err})
	}

	if err := db.Table("contact").Model(&contact).Updates(updateContactData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Failed to update contact", "data": err})
	}

	return c.JSON(contact)
}

func DeleteContact(c *fiber.Ctx) error {
	var contact models.Contact
	id := c.Params("id")
	db := database.DB
	db.Table("contact").Find(&contact, "id = ?", id)
	if contact.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"err": "Contact not found"})
	}
	if err := db.Table("contact").Delete(&contact).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"err": "Failed to delete contact", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted contact"})
}
