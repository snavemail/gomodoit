package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/snavemail/pomodoro-api/config"
	"github.com/snavemail/pomodoro-api/models"
)

func GetTasks(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var tasks []models.Task
	result := config.DB.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve tasks",
		})
	}

	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	
	if task.Title == "" || task.UserID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and UserID are required",
		})
	}
	
	if task.TaskType == models.TaskTypeTimed && task.TargetDuration == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Timed tasks require target duration",
		})
	}
	
	if task.TaskType == models.TaskTypeCountable && task.TargetRepetitions == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Countable tasks require target repetitions",
		})
	}
	
	result := config.DB.Create(&task)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create task",
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(task)
}

func GetTask(c *fiber.Ctx) error {
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}
	
	var task models.Task
	result := config.DB.First(&task, "task_id = ?", taskID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}
	
	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}
	
	var existingTask models.Task
	if result := config.DB.First(&existingTask, "task_id = ?", taskID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}
	
	var updatedTask models.Task
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	
	if updatedTask.TaskType == models.TaskTypeTimed && updatedTask.TargetDuration == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Timed tasks require target duration",
		})
	}
	
	if updatedTask.TaskType == models.TaskTypeCountable && updatedTask.TargetRepetitions == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Countable tasks require target repetitions",
		})
	}
	
	updatedTask.TaskID = taskID
	config.DB.Model(&existingTask).Updates(updatedTask)
	
	config.DB.First(&existingTask, "task_id = ?", taskID)
	
	return c.JSON(existingTask)
}

// DeleteTask deletes a task
func DeleteTask(c *fiber.Ctx) error {
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}
	
	var task models.Task
	if result := config.DB.First(&task, "task_id = ?", taskID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}
	
	config.DB.Delete(&task)
	
	return c.SendStatus(fiber.StatusNoContent)
}
