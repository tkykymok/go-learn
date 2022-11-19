package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null/v8"
	"time"
)

type Todo struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title     string    `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Completed bool      `boil:"completed" json:"completed" toml:"completed" yaml:"completed"`
	CreatedAt time.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
}

type TodoWithRelated struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title     string      `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Completed bool        `boil:"completed" json:"completed" toml:"completed" yaml:"completed"`
	Name      null.String `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
}

func GetTodoByIdResponse(data *Todo) *fiber.Map {
	todo := Todo{
		ID:        data.ID,
		Title:     data.Title,
		Completed: data.Completed,
		CreatedAt: data.CreatedAt,
	}
	return &fiber.Map{
		"data": todo,
	}
}

func GetAllTodosResponse(data *[]Todo) *fiber.Map {
	return &fiber.Map{
		"data": data,
	}
}

func GetTodosWithRelatedResponse(data *[]TodoWithRelated) *fiber.Map {
	return &fiber.Map{
		"data": data,
	}
}

func SuccessResponse(message string) *fiber.Map {
	return &fiber.Map{
		"data":    nil,
		"message": message,
	}
}
