package usecase

import (
	"context"
	"example.com/service-clean/module/todo/domain"
)

type useCaseImpl struct {
	repo domain.Repository
	createTodoUseCase
}

func (usecase useCaseImpl) CreateTodo(ctx context.Context, entity domain.TodoCreateInput) error {
	return usecase.createTodoUseCase.CreateTodo(ctx, entity)
}

func (usecase useCaseImpl) GetTodo(ctx context.Context, id string) (*domain.Todo, error) {
	return usecase.repo.GetTodo(ctx, id)
}

func (usecase useCaseImpl) GetTodos(ctx context.Context) ([]domain.Todo, error) {
	return usecase.repo.GetTodos(ctx)
}

func (usecase useCaseImpl) UpdateTodo(ctx context.Context, entity domain.Todo) error {
	return usecase.repo.UpdateTodo(ctx, entity)
}

func (usecase useCaseImpl) DeleteTodo(ctx context.Context, id string) error {
	return usecase.repo.DeleteTodo(ctx, id)
}

func New(repo domain.Repository) domain.UseCase {
	return &useCaseImpl{
		repo:              repo,
		createTodoUseCase: newCreateTodoUseCase(repo),
	}
}
