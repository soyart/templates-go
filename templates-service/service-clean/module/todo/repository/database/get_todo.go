package database

import (
	"context"
	"example.com/service-clean/module/todo/domain"
	"example.com/service-clean/module/todo/repository/database/model"
	"github.com/pkg/errors"
)

func (self adapterImpl) GetTodo(ctx context.Context, id string) (*domain.Todo, error) {
	query := &model.Todo{ID: id}
	getModel := &model.Todo{}

	err := self.db.First(ctx, query, getModel)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &domain.Todo{
		ID:          getModel.ID,
		Title:       getModel.Title,
		Status:      domain.Status(getModel.Status),
		Description: getModel.Description,
		TodoDate:    getModel.TodoDate,
		CreatedAt:   getModel.CreatedAt,
	}, nil
}
