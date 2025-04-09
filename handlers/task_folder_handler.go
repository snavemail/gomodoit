package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/snavemail/pomodoro-api/config"
	"github.com/snavemail/pomodoro-api/models"
)

func GetTaskFolders(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	
	var folders []models.TaskFolder
	result := config.DB.Where("user_id = ?", userID).Find(&folders)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve task folders",
		})
	}
	
	return c.JSON(folders)
}

func CreateTaskFolder(c *fiber.Ctx) error {
	folder := new(models.TaskFolder)
	
	if err := c.BodyParser(folder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	
	if folder.Name == "" || folder.UserID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name and UserID are required",
		})
	}
	
	result := config.DB.Create(&folder)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create task folder",
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(folder)
}

func UpdateTaskFolder(c *fiber.Ctx) error {
	folderID, err := uuid.Parse(c.Params("folderId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid folder ID",
		})
	}

	var existingFolder models.TaskFolder
	if result := config.DB.First(&existingFolder, "folder_id", folderID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task folder not found",
		})
	}

	var updateData struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	
	if updateData.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Folder name cannot be empty",
		})
	}
	
	config.DB.Model(&existingFolder).Update("name", updateData.Name)
	
	config.DB.First(&existingFolder, "folder_id = ?", folderID)
	
	return c.JSON(existingFolder)
}

func DeleteTaskFolder(c *fiber.Ctx) error {
	folderID, err := uuid.Parse(c.Params("folderId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid folder ID",
		})
	}
	
	var folder models.TaskFolder
	if result := config.DB.First(&folder, "folder_id = ?", folderID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Folder not found",
		})
	}
	
	var taskCount int64
	config.DB.Model(&models.Task{}).Where("folder_id = ?", folderID).Count(&taskCount)
	
	if taskCount > 0 {
		config.DB.Delete(&models.Task{}, "folder_id = ?", folderID)
	}
	
	config.DB.Delete(&folder)
	
	return c.SendStatus(fiber.StatusNoContent)
}
