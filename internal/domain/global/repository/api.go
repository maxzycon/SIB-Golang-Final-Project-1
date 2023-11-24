package repository

import (
	"context"

	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/model"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/util/pagination"
)

type GlobalRepository interface {
	// ---- Todos
	FindTodosById(ctx context.Context, id int) (resp *model.Todos, err error)
	FindAllTodos(ctx context.Context) (resp []*model.Todos, err error)
	FindTodosPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreateTodos(ctx context.Context, entity *model.Todos) (resp *int64, err error)
	UpdateTodosById(ctx context.Context, id int, entity *model.Todos) (resp *int64, err error)
	DeleteTodosById(ctx context.Context, id int) (resp *int64, err error)
}
