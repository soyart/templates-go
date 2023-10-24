package database

import (
	"context"
	"example.com/service-clean/module/todo/domain"
	"example.com/service-clean/module/todo/repository/database/model"
	"github.com/pkg/errors"
)

func (self adapterImpl) UpdateTodo(ctx context.Context, entity domain.Todo) error {
	query := &model.Todo{ID: entity.ID}
	updateModel := &model.Todo{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Status:      model.Status(entity.Status),
		TodoDate:    entity.TodoDate,
		CreatedAt:   entity.CreatedAt,
	}

	return errors.WithStack(self.db.Update(ctx, query, updateModel))
}
