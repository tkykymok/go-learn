package routes

import (
	"app/api/handlers"
	"app/pkg/todo"
	"github.com/gofiber/fiber/v2"
)

func TodoRouter(app fiber.Router, service todo.Service) {
	app.Get("/todos", handlers.GetAllTodos(service))
	app.Get("/todosWithRelated", handlers.GetTodosWithRelated(service))
	app.Get("/todo/:id", handlers.GetTodoById(service))
	app.Post("/todo", handlers.AddTodo(service))
	app.Put("/todo", handlers.UpdateTodo(service))
}
