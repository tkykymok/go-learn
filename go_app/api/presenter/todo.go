package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null/v8"
	"go_app/pkg/models"
)

type Todo struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title     null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Completed bool        `boil:"completed" json:"completed" toml:"completed" yaml:"completed"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
}

func TodoSuccessResponse(data *models.Todo) *fiber.Map {
	todo := Todo{
		ID:        data.ID,
		Title:     data.Title,
		Completed: data.Completed,
		CreatedAt: data.CreatedAt,
	}
	return &fiber.Map{
		"status": true,
		"data":   todo,
		"error":  nil,
	}
}

func TodosSuccessResponse(data *[]Todo) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}
