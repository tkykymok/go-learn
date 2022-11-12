package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go_app/api/presenter"
	"go_app/pkg/todo"
	"net/http"
)

func GetTodos(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create cancellable context.
		customContext, cancel := context.WithCancel(context.Background())
		defer cancel()

		fetched, err := service.FetchTodos(customContext)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.TodosSuccessResponse(fetched))
	}
}
