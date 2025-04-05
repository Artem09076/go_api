package controllers

import (
	"github.com/Artem09076/go_api.git/app/models"
	"github.com/Artem09076/go_api.git/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostGroup(c *fiber.Ctx) error {
	db := database.DB
	group := new(models.Group)
	if err := c.BodyParser(&group); err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Invalid credentials", "data": err})
	}
	group.ID = uuid.New()
	if err := db.Table("group").Create(&group).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create contact", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Created group", "data": group})
}

func GetGroup(c *fiber.Ctx) error {
	db := database.DB
	var groups []models.Group
	db.Table("group").Find(&groups)
	return c.JSON(groups)
}

func PutGroup(c *fiber.Ctx) error {
	var group models.Group
	id := c.Params("id")

	db := database.DB
	db.Table("group").Find(&group, "id = ?", id)
	if group.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"err": "Group not found"})
	}
	var updateGroupData map[string]any

	if err := c.BodyParser(&updateGroupData); err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Invalid credentials", "data": err})
	}
	if err := db.Table("group").Model(&group).Updates(updateGroupData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Failed to update group", "data": err})
	}
	return c.JSON(group)

}

func DeleteGroup(c *fiber.Ctx) error {
	var group models.Group
	id := c.Params("id")
	db := database.DB
	db.Table("group").Find(&group, "id = ?", id)
	if group.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"err": "Group not found"})
	}
	if err := db.Table("group").Delete(&group).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Failed to delete group", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted group"})
}
