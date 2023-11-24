package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/model"
)

func (s *GlobalService) GetAllTodos(ctx context.Context) (resp []*dto.TodosRow, err error) {
	data := make([]*model.Todos, 0)

	if err = s.db.WithContext(ctx).Model(&model.Todos{}).Find(&data).Error; err != nil {
		return
	}

	resp = make([]*dto.TodosRow, 0)

	for _, v := range data {
		resp = append(resp, &dto.TodosRow{
			ID:          v.ID,
			Status:      v.Status,
			Description: v.Description,
		})
	}

	return resp, nil
}

func (s *GlobalService) GetTodosById(ctx context.Context, id int) (resp *dto.TodosRow, err error) {
	row, err := s.globalRepository.FindTodosById(ctx, id)
	if err != nil {
		log.Errorf("err get Todos paginated")
		return
	}
	resp = &dto.TodosRow{
		ID:          row.ID,
		Status:      row.Status,
		Description: row.Description,
	}
	return
}

func (s *GlobalService) CreateTodos(ctx context.Context, payload *dto.PayloadTodos) (resp *int64, err error) {
	resp, err = s.globalRepository.CreateTodos(ctx, &model.Todos{
		Status:      payload.Status,
		Description: payload.Description,
	})
	if err != nil {
		log.Errorf("err create Todos")
		return
	}
	return
}

func (s *GlobalService) UpdateTodosById(ctx context.Context, id int, payload *dto.PayloadTodos) (resp *int64, err error) {
	resp, err = s.globalRepository.UpdateTodosById(ctx, id, &model.Todos{
		Status:      payload.Status,
		Description: payload.Description,
	})
	if err != nil {
		log.Errorf("err update Todos %d", id)
		return
	}
	return
}

func (s *GlobalService) DeleteTodosById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteTodosById(ctx, id)
	if err != nil {
		log.Errorf("err delete Todos %d", id)
		return
	}
	return
}
