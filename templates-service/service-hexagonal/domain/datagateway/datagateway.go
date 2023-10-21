package datagateway

import (
	"context"

	"example.com/servicehex/domain/entity"
)

type DataGatewayTodo interface {
	CreateTodo(context.Context, entity.Todo) error
	GetTodo(context.Context, string) (entity.Todo, error)
	UpdateTodo(context.Context, string, entity.Todo) (entity.Todo, error)
	DeleteTodo(context.Context, string) error
}
