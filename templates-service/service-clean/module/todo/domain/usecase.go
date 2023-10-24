package domain

import "context"

type UseCase interface {
	CreateTodo(ctx context.Context, todo TodoCreateInput) error
	GetTodo(context.Context, string) (*Todo, error)
	GetTodos(context.Context) ([]Todo, error)
	UpdateTodo(context.Context, Todo) error
	DeleteTodo(context.Context, string) error
}
