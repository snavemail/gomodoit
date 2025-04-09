package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snavemail/pomodoro-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	// User routes
	app.Post("/api/users", handlers.CreateUser)
	app.Get("/api/users/:userId", handlers.GetUser)
	
	// Task routes
	app.Get("/api/users/:userId/tasks", handlers.GetTasks)
	app.Post("/api/tasks", handlers.CreateTask)
	app.Get("/api/tasks/:taskId", handlers.GetTask)
	app.Put("/api/tasks/:taskId", handlers.UpdateTask)
	app.Delete("/api/tasks/:taskId", handlers.DeleteTask)
	
	// Task folder routes
	app.Get("/api/users/:userId/folders", handlers.GetTaskFolders)
	app.Post("/api/folders", handlers.CreateTaskFolder)
	app.Put("/api/folders/:folderId", handlers.UpdateTaskFolder) 
	app.Delete("/api/folders/:folderId", handlers.DeleteTaskFolder)  
	
	// Tag routes
	app.Get("/api/users/:userId/tags", handlers.GetTags)
	app.Post("/api/tags", handlers.CreateTag)
	app.Get("/api/tags/:tagId", handlers.GetTag)
	app.Put("/api/tags/:tagId", handlers.UpdateTag)
	app.Delete("/api/tags/:tagId", handlers.DeleteTag)
	
	// Pomodoro session routes
	app.Get("/api/users/:userId/pomodoro-sessions", handlers.GetPomodoroSessions)
	app.Post("/api/pomodoro-sessions", handlers.CreatePomodoroSession)
	app.Get("/api/pomodoro-sessions/:sessionId", handlers.GetPomodoroSession)
	
	// Work session routes
	app.Get("/api/pomodoro-sessions/:sessionId/work-sessions", handlers.GetWorkSessions)
	app.Post("/api/work-sessions", handlers.CreateWorkSession)
	
	// Break session routes
	app.Get("/api/pomodoro-sessions/:sessionId/break-sessions", handlers.GetBreakSessions)
	app.Post("/api/break-sessions", handlers.CreateBreakSession)
}
