package todo

import (
	"app/api/presenter"
	"app/api/requests"
	"app/pkg/models"
	"context"
)

type Service interface {
	FetchAllTodos(ctx context.Context) (*[]presenter.Todo, error)
	FetchTodosWithRelated(ctx context.Context) (*[]presenter.TodoWithRelated, error)
	FetchTodoById(ctx context.Context, id int) (*presenter.Todo, error)
	InsertTodo(ctx context.Context, todo *requests.AddTodo) error
	UpdateTodo(ctx context.Context, todo *requests.UpdateTodo) error
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

func (s service) FetchTodosWithRelated(ctx context.Context) (*[]presenter.TodoWithRelated, error) {
	return s.repository.ReadTodosWithRelated(ctx)
}

func (s service) FetchTodoById(ctx context.Context, id int) (*presenter.Todo, error) {
	return s.repository.ReadTodoById(ctx, id)
}

func (s service) InsertTodo(ctx context.Context, todo *requests.AddTodo) error {
	cTodo := models.Todo{
		Title: todo.Title,
	}
	return s.repository.CreateTodo(ctx, &cTodo)
}

func (s service) UpdateTodo(ctx context.Context, todo *requests.UpdateTodo) error {
	uTodo := models.Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	}
	return s.repository.UpdateTodo(ctx, &uTodo)
}
