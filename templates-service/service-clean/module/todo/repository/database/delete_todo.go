package database

import (
	"context"
	"example.com/service-clean/module/todo/repository/database/model"
	"github.com/pkg/errors"
)

func (self adapterImpl) DeleteTodo(ctx context.Context, id string) error {
	query := &model.Todo{ID: id}
	return errors.WithStack(self.db.Delete(ctx, query))
}
