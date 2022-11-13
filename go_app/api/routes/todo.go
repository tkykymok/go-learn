package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go_app/api/handlers"
	"go_app/pkg/todo"
)

func TodoRouter(app fiber.Router, service todo.Service, validate *validator.Validate) {
	app.Get("/todos", handlers.GetAllTodos(service))
	app.Get("/todo/:id", handlers.GetTodoById(service))
	app.Post("/todo", handlers.AddTodo(service, validate))
	app.Put("/todo", handlers.UpdateTodo(service))
}
