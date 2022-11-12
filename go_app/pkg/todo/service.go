package todo

import (
	"context"
	"go_app/api/presenter"
)

type Service interface {
	FetchAllTodos(ctx context.Context) (*[]presenter.Todo, error)
	FetchTodoById(ctx context.Context, id int) (*presenter.Todo, error)
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
