package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/snavemail/pomodoro-api/config"
	"github.com/snavemail/pomodoro-api/models"
)

func GetPomodoroSessions(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	
	var sessions []models.PomodoroSession
	result := config.DB.Where("user_id = ?", userID).Find(&sessions)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve pomodoro sessions",
		})
	}
	
	return c.JSON(sessions)
}

func CreatePomodoroSession(c *fiber.Ctx) error {
	session := new(models.PomodoroSession)
	
	if err := c.BodyParser(session); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	
	if session.UserID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UserID is required",
		})
	}
	
	result := config.DB.Create(&session)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create pomodoro session",
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(session)
}

func GetPomodoroSession(c *fiber.Ctx) error {
	sessionID, err := uuid.Parse(c.Params("sessionId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}
	
	var session models.PomodoroSession
	result := config.DB.Preload("WorkSessions").Preload("BreakSessions").First(&session, "session_id = ?", sessionID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Pomodoro session not found",
		})
	}
	
	return c.JSON(session)
}
