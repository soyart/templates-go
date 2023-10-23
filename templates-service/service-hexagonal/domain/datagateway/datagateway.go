package datagateway

import (
	"context"

	"example.com/servicehex/domain/entity"
)

type DataGatewayTodo interface {
	CreateTodo(ctx context.Context, todo entity.Todo) error
	GetTodo(ctx context.Context, userID string, todoID string) (entity.Todo, error)
	GetTodos(ctx context.Context, userID string) ([]entity.Todo, error)
	UpdateTodo(ctx context.Context, userID string, todoID string, update entity.Todo) error
	DeleteTodo(ctx context.Context, userID string, todoID string) error
	DeleteTodos(ctx context.Context, userID string) error
}

type DataGatewayUser interface {
	CreateUser(context.Context, entity.User) error
	GetUser(context.Context, string) (entity.User, error)
	GetUserByUsername(context.Context, string) (entity.User, error)
	UpdateUser(context.Context, string, entity.User) error
	DeleteUser(context.Context, string) error
}

// type dataGateway[K comparable, T any] interface {
// 	Create(context.Context, T) error
// 	GetOne(context.Context, K) (T, error)
// 	GetAll(context.Context) ([]T, error)
// 	Update(context.Context, K, T) error
// 	Delete(context.Context, K) error
// }

// type (
// 	DataGatewayUser dataGateway[string, entity.Todo]
// 	DataGatewayTodo dataGateway[string, entity.User]
// )
