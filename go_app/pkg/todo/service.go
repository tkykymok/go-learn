package todo

import (
	"context"
	"go_app/api/presenter"
	"go_app/api/requests"
	"go_app/pkg/models"
)

type Service interface {
	FetchAllTodos(ctx context.Context) (*[]presenter.Todo, error)
	FetchTodoById(ctx context.Context, id int) (*presenter.Todo, error)
	InsertTodo(ctx context.Context, todo requests.AddTodo) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) FetchAllTodos(ctx context.Context) (*[]presenter.Todo, error) {
	return s.repository.ReadAllTodos(ctx)
}

func (s service) FetchTodoById(ctx context.Context, id int) (*presenter.Todo, error) {
	return s.repository.ReadTodoById(ctx, id)
}

func (s service) InsertTodo(ctx context.Context, todo requests.AddTodo) error {
	cTodo := models.Todo{
		Title: todo.Title,
	}

	return s.repository.CreateTodo(ctx, &cTodo)
}
