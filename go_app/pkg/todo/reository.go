package todo

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go_app/api/presenter"
	"go_app/pkg/models"
)

type Repository interface {
	ReadAllTodos(ctx context.Context) (*[]presenter.Todo, error)
	ReadTodoById(ctx context.Context, id int) (*presenter.Todo, error)
	CreateTodo(ctx context.Context, todo *models.Todo)
}

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}

func (r repository) ReadAllTodos(ctx context.Context) (*[]presenter.Todo, error) {
	var todos []presenter.Todo
	result, err := models.Todos().AllG(ctx)
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

func (r repository) ReadTodoById(ctx context.Context, id int) (*presenter.Todo, error) {
	result, err := models.FindTodoG(ctx, id)
	if err != nil {
		return nil, err
	}

	todo := presenter.Todo{
		ID:        result.ID,
		Title:     result.Title,
		Completed: result.Completed,
		CreatedAt: result.CreatedAt,
	}

	return &todo, nil
}

func (r repository) CreateTodo(ctx context.Context, todo *models.Todo) {
	err := todo.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return
	}
}
