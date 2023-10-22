package datagateway

import (
	"context"

	"example.com/servicehex/domain/entity"
)

type DataGatewayTodo interface {
	CreateTodo(context.Context, entity.Todo) error
	GetTodo(context.Context, string) (entity.Todo, error)
	GetTodos(context.Context) ([]entity.Todo, error)
	UpdateTodo(context.Context, string, entity.Todo) error
	DeleteTodo(context.Context, string) error
}

type DataGatewayUser interface {
	CreateUser(context.Context, entity.User) error
	GetUser(context.Context, string) (entity.User, error)
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
