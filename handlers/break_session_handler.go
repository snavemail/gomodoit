package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/snavemail/pomodoro-api/config"
	"github.com/snavemail/pomodoro-api/models"
)


func GetBreakSessions(c *fiber.Ctx) error {
	sessionID, err := uuid.Parse(c.Params("sessionId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}
	
	var breakSessions []models.BreakSession
	result := config.DB.Preload("Task").Where("session_id = ?", sessionID).Find(&breakSessions)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve break sessions",
		})
	}
	
	return c.JSON(breakSessions)
}

func CreateBreakSession(c *fiber.Ctx) error {
	breakSession := new(models.BreakSession)
	
	if err := c.BodyParser(breakSession); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	
	if breakSession.SessionID == uuid.Nil || breakSession.Status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "SessionID and Status are required",
		})
	}
	
	if breakSession.Status != models.StatusCompleted && 
	   breakSession.Status != models.StatusSkipped {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Status must be 'completed' or 'skipped'",
		})
	}
	
	if breakSession.CompletionType == models.MetricDuration && breakSession.ActualDuration == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Duration completion type requires actual duration",
		})
	}
	
	if breakSession.CompletionType == models.MetricRepetitions && breakSession.ActualRepetitions == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Repetitions completion type requires actual repetitions",
		})
	}
	
	var session models.PomodoroSession
	if result := config.DB.First(&session, "session_id = ?", breakSession.SessionID); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid pomodoro session ID",
		})
	}
	
	if breakSession.TaskID != uuid.Nil {
		var task models.Task
		if result := config.DB.First(&task, "task_id = ?", breakSession.TaskID); result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid task ID",
			})
		}
	}
	
	result := config.DB.Create(&breakSession)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create break session",
		})
	}
	
	if breakSession.TaskID != uuid.Nil {
		config.DB.Preload("Task").First(&breakSession, "break_id = ?", breakSession.BreakID)
	}
	
	return c.Status(fiber.StatusCreated).JSON(breakSession)
}
