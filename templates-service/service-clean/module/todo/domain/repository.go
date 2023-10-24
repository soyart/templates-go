package domain

import "context"

type Repository interface {
	CreateTodo(context.Context, Todo) error
	GetTodo(context.Context, string) (*Todo, error)
	GetTodos(context.Context) ([]Todo, error)
	UpdateTodo(context.Context, Todo) error
	DeleteTodo(context.Context, string) error
}
