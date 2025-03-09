package routes

import (
	"github.com/gofiber/fiber/v2"
	"tasklist/internal/handler"
)

func TaskListRoutes(app *fiber.App, handler *handler.TaskHandler) {
	api := app.Group("/tasks")

	api.Post("", handler.CreateTaskHandler)
	api.Get("", handler.GetAllTasksHandler)
	api.Put("/:id", handler.UpdateTaskHandler)
	api.Delete("/:id", handler.DeleteTaskHandler)
}
