package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/snavemail/pomodoro-api/config"
	"github.com/snavemail/pomodoro-api/models"
)


func GetWorkSessions(c *fiber.Ctx) error {
	sessionID, err := uuid.Parse(c.Params("sessionId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}
	
	var workSessions []models.WorkSession
	result := config.DB.Preload("Tag").Where("session_id = ?", sessionID).Find(&workSessions)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve work sessions",
		})
	}
	
	return c.JSON(workSessions)
}

func CreateWorkSession(c *fiber.Ctx) error {
	workSession := new(models.WorkSession)
	
	if err := c.BodyParser(workSession); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	
	if workSession.SessionID == uuid.Nil || workSession.Duration <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "SessionID and Duration are required",
		})
	}
	
	var session models.PomodoroSession
	if result := config.DB.First(&session, "session_id = ?", workSession.SessionID); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid pomodoro session ID",
		})
	}
	
	result := config.DB.Create(&workSession)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create work session",
		})
	}
	
	if workSession.TagID != uuid.Nil {
		config.DB.Preload("Tag").First(&workSession, "work_id = ?", workSession.WorkID)
	}
	
	return c.Status(fiber.StatusCreated).JSON(workSession)
}
