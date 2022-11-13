package routes

import (
	"app/api/handlers"
	"app/pkg/todo"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func TodoRouter(app fiber.Router, service todo.Service, validate *validator.Validate) {
	app.Get("/todos", handlers.GetAllTodos(service))
	app.Get("/todo/:id", handlers.GetTodoById(service))
	app.Post("/todo", handlers.AddTodo(service, validate))
	app.Put("/todo", handlers.UpdateTodo(service))
}
