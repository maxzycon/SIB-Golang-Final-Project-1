package service

import (
	"context"

	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/global/dto"
)

type GlobalService interface {
	// ---- Todos
	GetAllTodos(ctx context.Context) (resp []*dto.TodosRow, err error)
	GetTodosById(ctx context.Context, id int) (resp *dto.TodosRow, err error)
	CreateTodos(ctx context.Context, payload *dto.PayloadTodos) (resp *int64, err error)
	UpdateTodosById(ctx context.Context, id int, payload *dto.PayloadTodos) (resp *int64, err error)
	DeleteTodosById(ctx context.Context, id int) (resp *int64, err error)
}
