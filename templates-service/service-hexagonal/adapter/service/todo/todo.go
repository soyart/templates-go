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

func getTodoByID(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userID string,
	todoID string,
) (
	entity.Todo,
	error,
) {
	todo, err := repo.GetTodo(ctx, userID, todoID)
	if err != nil {
		return todo, errors.WithStack(err)
	}

	if todo.UserID != userID {
		return todo, core.WrongUserID(todoID, userID) //nolint:wrapcheck
	}

	return todo, nil
}

func getTodos(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userID string,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	todos, err := repo.GetTodos(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	core.SortTodos(todos, sortMode)

	return todos, nil
}

func getExpiredTodos(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userID string,
	cutoff time.Time,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	todos, err := repo.GetTodos(ctx, userID)
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
	userID string,
	pattern string,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	todos, err := repo.GetTodos(ctx, userID)
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

func expireTodoByID(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userID string,
	todoID string,
) error {
	todo, err := repo.GetTodo(ctx, userID, todoID)
	if err != nil {
		return errors.WithStack(err)
	}
	if todo.UserID != userID {
		return core.WrongUserID(todoID, userID) //nolint:wrapcheck
	}

	if todo.Expired {
		return nil
	}

	core.ExpireTodo(&todo)

	if err := repo.UpdateTodo(ctx, userID, todoID, todo); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func deleteTodoByID(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userID string,
	todoID string,
) error {
	todo, err := repo.GetTodo(ctx, userID, todoID)
	if err != nil {
		return errors.WithStack(err)
	}
	if todo.UserID != userID {
		return core.WrongUserID(todoID, userID) //nolint:wrapcheck
	}

	if err := repo.DeleteTodo(ctx, userID, todoID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func deleteTodos(
	ctx context.Context,
	repo datagateway.DataGatewayTodo,
	userID string,
) error {
	if err := repo.DeleteTodos(ctx, userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
