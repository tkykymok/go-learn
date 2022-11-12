package todo

import (
	"context"
	"go_app/api/presenter"
)

type Service interface {
	FetchTodos(ctx context.Context) (*[]presenter.Todo, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) FetchTodos(ctx context.Context) (*[]presenter.Todo, error) {
	return s.repository.ReadTodos(ctx)
}
