package todo

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"

	"example.com/servicehex/domain/core"
	"example.com/servicehex/domain/datagateway"
	"example.com/servicehex/domain/entity"
)

func createTodo(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	todo entity.Todo,
) error {
	err := repo.CreateTodo(ctx, todo)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func getTodoById(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userId string,
	todoId string,
) (
	entity.Todo,
	error,
) {
	todo, err := repo.GetTodo(ctx, userId, todoId)
	if err != nil {
		return todo, errors.WithStack(err)
	}

	if todo.UserId != userId {
		return todo, core.WrongUserId(todoId, userId)
	}

	return todo, nil
}

func getTodos(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userId string,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	todos, err := repo.GetTodos(ctx, userId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	core.SortTodos(todos, sortMode)

	return todos, nil
}

func getExpiredTodos(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userId string,
	cutoff time.Time,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	todos, err := repo.GetTodos(ctx, userId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var expired []entity.Todo
	for i := range todos {
		t := &todos[i]

		switch {
		case t.Expired, t.Deadline.Before(cutoff):
			expired = append(expired, *t)
		}
	}

	core.SortTodos(expired, sortMode)

	return expired, nil
}

func matchTodoTextPattern(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userId string,
	pattern string,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	todos, err := repo.GetTodos(ctx, userId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var matched []entity.Todo
	for i := range todos {
		t := &todos[i]
		if strings.Contains(t.Text, pattern) {
			matched = append(matched, *t)
		}
	}

	core.SortTodos(matched, sortMode)

	return matched, nil
}

func expireTodoById(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userId string,
	todoId string,
) error {
	todo, err := repo.GetTodo(ctx, userId, todoId)
	if err != nil {
		return errors.WithStack(err)
	}
	if todo.UserId != userId {
		return core.WrongUserId(todoId, userId)
	}

	if todo.Expired {
		return nil
	}

	core.ExpireTodo(&todo)

	if err := repo.UpdateTodo(ctx, userId, todoId, todo); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func deleteTodoById(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userId string,
	todoId string,
) error {
	todo, err := repo.GetTodo(ctx, userId, todoId)
	if err != nil {
		return errors.WithStack(err)
	}
	if todo.UserId != userId {
		return core.WrongUserId(todoId, userId)
	}

	if err := repo.DeleteTodo(ctx, userId, todoId); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func deleteTodos(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userId string,
) error {
	if err := repo.DeleteTodos(ctx, userId); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
