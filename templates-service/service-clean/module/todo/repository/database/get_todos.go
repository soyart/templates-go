package database

import (
	"context"
	"example.com/service-clean/module/todo/domain"
	"example.com/service-clean/module/todo/repository/database/model"
	"github.com/pkg/errors"
)

func (self adapterImpl) GetTodos(ctx context.Context) ([]domain.Todo, error) {
	models := []model.Todo{}

	err := self.db.Find(ctx, nil, models)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	results := make([]domain.Todo, len(models))
	for i := range results {
		results[i] = domain.Todo{
			ID:          models[i].ID,
			Title:       models[i].Title,
			Status:      domain.Status(models[i].Status),
			Description: models[i].Description,
			TodoDate:    models[i].TodoDate,
			CreatedAt:   models[i].CreatedAt,
		}
	}

	return results, nil
}
