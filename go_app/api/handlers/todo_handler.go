package handlers

import (
	"app/api/message"
	"app/api/presenter"
	"app/api/requests"
	"app/api/validation"
	"app/pkg/todo"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetAllTodos(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create cancellable context.
		customContext, cancel := context.WithCancel(context.Background())
		defer cancel()

		fetched, err := service.FetchAllTodos(customContext)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.GetAllTodosResponse(fetched))
	}
}

func GetTodoById(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customContext, cancel := context.WithCancel(context.Background())
		defer cancel()

		id, _ := c.ParamsInt("id", 0)

		fetched, err := service.FetchTodoById(customContext, id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.GetTodoByIdResponse(fetched))
	}
}

func AddTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customContext, cancel := context.WithCancel(context.Background())
		defer cancel()

		var request requests.AddTodo
		err := c.BodyParser(&request)
		if err != nil {
			return err
		}

		// バリデーションチェック
		err = validation.ValidateStruct(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ValidationErrorResponse(err))
		}

		err = service.InsertTodo(customContext, &request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err))
		}

		return c.JSON(presenter.SuccessResponse(message.GetMessage("success", "登録")))
	}
}

func UpdateTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customContext, cancel := context.WithCancel(context.Background())
		defer cancel()

		var request requests.UpdateTodo
		err := c.BodyParser(&request)
		if err != nil {
			return err
		}

		err = service.UpdateTodo(customContext, &request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err))
		}

		return c.JSON(presenter.SuccessResponse("更新が完了しました。"))
	}
}
