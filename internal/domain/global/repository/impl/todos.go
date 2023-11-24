package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/model"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/util/pagination"
)

func (r *GlobalRepository) FindTodosById(ctx context.Context, id int) (resp *model.Todos, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllTodos(ctx context.Context) (resp []*model.Todos, err error) {
	resp = make([]*model.Todos, 0)
	tx := r.db.WithContext(ctx).Model(&model.Todos{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindTodosPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Todos []*model.Todos = make([]*model.Todos, 0)
	sql := r.db.WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Todos, &resp.Paginator, sql)).Find(&Todos)
	resp.Items = Todos
	return
}

func (r *GlobalRepository) CreateTodos(ctx context.Context, entity *model.Todos) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Todos{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateTodosById(ctx context.Context, id int, entity *model.Todos) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteTodosById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Todos{}, id)
	return &tx.RowsAffected, tx.Error
}
