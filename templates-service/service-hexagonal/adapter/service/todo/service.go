package todo

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"example.com/servicehex/domain/core"
	"example.com/servicehex/domain/datagateway"
	"example.com/servicehex/domain/entity"
)

type service struct {
	repo datagateway.DataGatewayTodo
}

func New(repo datagateway.DataGatewayTodo) core.PortTodo {
	return &service{repo: repo}
}

func (s *service) CreateTodo(ctx context.Context, todo entity.Todo) error {
	if err := createTodo(ctx, s.repo, todo); err != nil {
		return errors.Wrap(err, "failed to create new todo")
	}

	return nil
}

func (s *service) GetTodoByID(
	ctx context.Context,
	userID string,
	todoID string,
) (
	entity.Todo,
	error,
) {
	todo, err := getTodoByID(ctx, s.repo, userID, todoID)
	if err != nil {
		return todo, errors.Wrap(err, "failed to get user todo")
	}

	return todo, nil
}

func (s *service) GetTodos(
	ctx context.Context,
	userID string,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	todos, err := getTodos(ctx, s.repo, userID, sortMode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user todos")
	}

	return todos, nil
}

func (s *service) GetExpiredTodos(
	ctx context.Context,
	userID string,
	cutoff time.Time,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	expireds, err := getExpiredTodos(ctx, s.repo, userID, cutoff, sortMode) //nolint:misspell
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get expired todo for user %s at %s cutoff", userID, cutoff)
	}

	return expireds, nil //nolint:misspell
}

func (s *service) MatchTodoTextPattern(
	ctx context.Context,
	userID string,
	pattern string,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	matched, err := matchTodoTextPattern(ctx, s.repo, userID, pattern, sortMode)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to match todos for user %s with %s pattern", userID, pattern)
	}

	return matched, nil
}

func (s *service) ExpireTodoByID(
	ctx context.Context,
	userID string,
	todoID string,
) error {
	if err := expireTodoByID(ctx, s.repo, userID, todoID); err != nil {
		return errors.Wrapf(err, "failed to expire todo %s for user %s", todoID, userID)
	}

	return nil
}

func (s *service) DeleteTodoByID(
	ctx context.Context,
	userID string,
	todoID string,
) error {
	if err := deleteTodoByID(ctx, s.repo, userID, todoID); err != nil {
		return errors.Wrapf(err, "failed to delete todo %s for user %s", todoID, userID)
	}

	return nil
}

func (s *service) DeleteTodos(ctx context.Context, userID string) error {
	if err := deleteTodos(ctx, s.repo, userID); err != nil {
		return errors.Wrapf(err, "failed to delete todos for user %s", userID)
	}

	return nil
}
