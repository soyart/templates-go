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

func (s *service) GetTodoById(
	ctx context.Context,
	userId string,
	todoId string,
) (
	entity.Todo,
	error,
) {
	todo, err := getTodoById(ctx, s.repo, userId, todoId)
	if err != nil {
		return todo, errors.Wrap(err, "failed to get user todo")
	}

	return todo, nil
}

func (s *service) GetTodos(
	ctx context.Context,
	userId string,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	todos, err := getTodos(ctx, s.repo, userId, sortMode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user todos")
	}

	return todos, nil
}

func (s *service) GetExpiredTodos(
	ctx context.Context,
	userId string,
	cutoff time.Time,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	expireds, err := getExpiredTodos(ctx, s.repo, userId, cutoff, sortMode)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get expired todo for user %s at %s cutoff", userId, cutoff)
	}

	return expireds, nil
}

func (s *service) MatchTodoTextPattern(
	ctx context.Context,
	userId string,
	pattern string,
	sortMode core.EnumSortMode,
) (
	[]entity.Todo,
	error,
) {
	matched, err := matchTodoTextPattern(ctx, s.repo, userId, pattern, sortMode)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to match todos for user %s with %s pattern", userId, pattern)
	}

	return matched, nil
}

func (s *service) ExpireTodoById(
	ctx context.Context,
	userId string,
	todoId string,
) error {
	if err := expireTodoById(ctx, s.repo, userId, todoId); err != nil {
		return errors.Wrapf(err, "failed to expire todo %s for user %s", todoId, userId)
	}

	return nil
}

func (s *service) DeleteTodoById(
	ctx context.Context,
	userId string,
	todoId string,
) error {
	if err := deleteTodoById(ctx, s.repo, userId, todoId); err != nil {
		return errors.Wrapf(err, "failed to delete todo %s for user %s", todoId, userId)
	}

	return nil
}

func (s *service) DeleteTodos(ctx context.Context, userId string) error {
	if err := deleteTodos(ctx, s.repo, userId); err != nil {
		return errors.Wrapf(err, "failed to delete todos for user %s", userId)
	}

	return nil
}
