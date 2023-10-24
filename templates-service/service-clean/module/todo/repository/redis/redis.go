package redis

import (
	"context"
	"example.com/service-clean/module/todo/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type Repository interface {
	CreateTodo(context.Context, domain.Todo) error
	GetTodo(context.Context, string) (*domain.Todo, error)
	GetTodos(context.Context) ([]domain.Todo, error)
	UpdateTodo(context.Context, domain.Todo) error
	DeleteTodo(context.Context, string) error
}

type repositoryImpl struct {
	rd Wrapper
}

func (self repositoryImpl) CreateTodo(ctx context.Context, todo domain.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (self repositoryImpl) GetTodo(ctx context.Context, id string) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (self repositoryImpl) GetTodos(ctx context.Context) ([]domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (self repositoryImpl) UpdateTodo(ctx context.Context, todo domain.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (self repositoryImpl) DeleteTodo(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func New(rd Wrapper) Repository {
	return &repositoryImpl{rd: rd}
}

type Wrapper interface {
	Set(context.Context, string, interface{}, time.Duration) *redis.StatusCmd
}
