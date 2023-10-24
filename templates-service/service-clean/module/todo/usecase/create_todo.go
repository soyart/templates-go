package usecase

import (
	"context"
	"example.com/service-clean/module/todo/common"
	"example.com/service-clean/module/todo/domain"
	"github.com/google/uuid"
	"time"
)

type createTodoUseCase interface {
	CreateTodo(context.Context, domain.TodoCreateInput) error
}

type createTodoImpl struct {
	repo domain.Repository
}

func newCreateTodoUseCase(repo domain.Repository) createTodoUseCase {
	return &createTodoImpl{repo: repo}
}

func (self createTodoImpl) CreateTodo(ctx context.Context, entity domain.TodoCreateInput) error {
	if err := self.validateInput(entity); err != nil {
		return self.wrapError(err)
	}

	if err := self.repo.CreateTodo(ctx, self.prepare(entity)); err != nil {
		return self.wrapError(err)
	}

	return nil
}

func (self createTodoImpl) prepare(input domain.TodoCreateInput) domain.Todo {
	return domain.Todo{
		ID:          uuid.NewString(),
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		TodoDate:    input.TodoDate,
		CreatedAt:   time.Now(),
	}
}

func (self createTodoImpl) validateInput(input domain.TodoCreateInput) error {
	return common.ValidateStruct(input)
}

func (self createTodoImpl) wrapError(err error) error {
	return common.WrapError(common.ErrCreateTodo.Error()+` use case`, err)
}
