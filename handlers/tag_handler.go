package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/snavemail/pomodoro-api/config"
	"github.com/snavemail/pomodoro-api/models"
)

func GetTags(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	
	var tags []models.Tag
	result := config.DB.Where("user_id = ?", userID).Find(&tags)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve tags",
		})
	}
	
	return c.JSON(tags)
}

func CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)
	
	if err := c.BodyParser(tag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	
	if tag.Name == "" || tag.UserID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name and UserID are required",
		})
	}
	
	result := config.DB.Create(&tag)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create tag",
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(tag)
}

func GetTag(c *fiber.Ctx) error {
	tagID, err := uuid.Parse(c.Params("tagId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid tag ID",
		})
	}
	
	var tag models.Tag
	result := config.DB.First(&tag, "tag_id = ?", tagID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tag not found",
		})
	}
	
	return c.JSON(tag)
}

func UpdateTag(c *fiber.Ctx) error {
	tagID, err := uuid.Parse(c.Params("tagId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid tag ID",
		})
	}
	
	var existingTag models.Tag
	if result := config.DB.First(&existingTag, "tag_id = ?", tagID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tag not found",
		})
	}
	
	var updatedTag models.Tag
	if err := c.BodyParser(&updatedTag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	
	updatedTag.TagID = tagID 
	config.DB.Model(&existingTag).Updates(updatedTag)
	
	config.DB.First(&existingTag, "tag_id = ?", tagID)
	
	return c.JSON(existingTag)
}

func DeleteTag(c *fiber.Ctx) error {
	tagID, err := uuid.Parse(c.Params("tagId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid tag ID",
		})
	}
	
	var tag models.Tag
	if result := config.DB.First(&tag, "tag_id = ?", tagID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tag not found",
		})
	}
	
	var count int64
	config.DB.Model(&models.WorkSession{}).Where("tag_id = ?", tagID).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot delete tag as it is used in work sessions",
		})
	}
	
	config.DB.Delete(&tag)
	
	return c.SendStatus(fiber.StatusNoContent)
}
