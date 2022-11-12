package todo

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go_app/api/presenter"
	"go_app/pkg/models"
)

type Repository interface {
	ReadTodos(ctx context.Context) (*[]presenter.Todo, error)
}

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}

func (r repository) ReadTodos(ctx context.Context) (*[]presenter.Todo, error) {
	var todos []presenter.Todo

	result, err := models.Todos().All(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	for _, t := range result {
		todo := presenter.Todo{
			ID:        t.ID,
			Title:     t.Title,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		}
		todos = append(todos, todo)
	}

	return &todos, nil
}
