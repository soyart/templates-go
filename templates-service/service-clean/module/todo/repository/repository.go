package repository

import (
	"context"
	"example.com/service-clean/module/todo/domain"
	"example.com/service-clean/module/todo/repository/database"
	redisrepo "example.com/service-clean/module/todo/repository/redis"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTodo(ctx context.Context, todo domain.Todo) error
	GetTodo(context.Context, string) (*domain.Todo, error)
	GetTodos(context.Context) ([]domain.Todo, error)
	UpdateTodo(context.Context, domain.Todo) error
	DeleteTodo(context.Context, string) error
}

func New(repoClient interface{}) Repository {
	switch repoClient.(type) {
	case *redis.Client:
		return redisrepo.New(repoClient.(*redis.Client))
	case *gorm.DB:
		return database.New(repoClient.(*gorm.DB))
	default:
		panic("invalid repository client")
	}
}
