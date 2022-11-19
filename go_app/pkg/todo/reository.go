package todo

import (
	"app/api/presenter"
	"app/api/requests"
	"app/pkg/exmodels"
	"app/pkg/models"
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"strings"
)

type Repository interface {
	ReadAllTodos(ctx context.Context) (*[]presenter.Todo, error)
	ReadTodosWithRelated(ctx context.Context, request *requests.GetTodosWithRelated) (*[]presenter.TodoWithRelated, error)
	ReadTodoById(ctx context.Context, id int) (*presenter.Todo, error)
	CreateTodo(ctx context.Context, todo *models.Todo) error
	UpdateTodo(ctx context.Context, todo *models.Todo) error
}

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}

func (r repository) ReadAllTodos(ctx context.Context) (*[]presenter.Todo, error) {
	todos := make([]presenter.Todo, 0)
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

func (r repository) ReadTodosWithRelated(ctx context.Context, request *requests.GetTodosWithRelated) (*[]presenter.TodoWithRelated, error) {
	todos := make([]presenter.TodoWithRelated, 0)

	// SELECTするカラム
	selectCols := []string{
		models.TodoTableColumns.ID,
		models.TodoTableColumns.Title,
		models.TodoTableColumns.Completed,
		models.UserTableColumns.Name,
		models.TodoTableColumns.CreatedAt,
	}

	// QueryModの生成
	mods := []qm.QueryMod{
		qm.Select(strings.Join(selectCols, ",")),
		qm.LeftOuterJoin("users on todos.userId = users.id"),
	}
	// WHERE句
	if request.ID != 0 {
		mods = append(mods, qm.Where("todos.id=?", request.ID))
	}
	if request.UserId != 0 {
		mods = append(mods, qm.Where("users.id=?", request.UserId))
	}

	var result []exmodels.TodoWithRelated
	err := models.Todos(mods...).BindG(ctx, &result)
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range result {
		todo := presenter.TodoWithRelated{
			ID:        t.ID,
			Title:     t.Title,
			Completed: t.Completed,
			Name:      t.Name,
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

func (r repository) CreateTodo(ctx context.Context, todo *models.Todo) error {
	err := todo.InsertG(ctx, boil.Infer())
	if err != nil {
		return nil
	}
	return err
}

func (r repository) UpdateTodo(ctx context.Context, todo *models.Todo) error {
	_, err := todo.UpdateG(ctx, boil.Infer())
	if err != nil {
		return nil
	}
	return err
}
