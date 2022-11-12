package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_app/api/handlers"
	"go_app/pkg/todo"
)

func TodoRouter(app fiber.Router, service todo.Service) {
	app.Get("/todos", handlers.GetTodos(service))
}
