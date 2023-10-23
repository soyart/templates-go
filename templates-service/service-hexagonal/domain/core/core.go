package core

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"example.com/servicehex/domain/entity"
)

const MinPasswordLen = 6

type PortTodo interface {
	CreateTodo(ctx context.Context, todo entity.Todo) error
	GetTodoById(ctx context.Context, userId, todoId string) (entity.Todo, error)
	GetTodos(ctx context.Context, userId string, sortMode EnumSortMode) ([]entity.Todo, error)
	GetExpiredTodos(ctx context.Context, userId string, cutoff time.Time, sortMode EnumSortMode) ([]entity.Todo, error)
	MatchTodoTextPattern(ctx context.Context, userId, pattern string, sortMode EnumSortMode) ([]entity.Todo, error)
	ExpireTodoById(ctx context.Context, userId, todoId string) error
	DeleteTodoById(ctx context.Context, userId, todoId string) error
	DeleteTodos(ctx context.Context, userId string) error
}

type PortUser interface {
	Register(ctx context.Context, user entity.User) error
	Login(ctx context.Context, username string, password []byte) (entity.User, error)
	ChangePassword(ctx context.Context, userId string, password []byte, newPassword []byte) error
	DeleteUser(ctx context.Context, userId string) error
}

func ExpireTodo(todo *entity.Todo) {
	todo.Expired, todo.ExpiredAt = true, time.Now()
}

func TodoId() string {
	return randUuid()
}

func UserId(username string) string {
	return uuid.New().String()
}

func ValidatePassword(p []byte) error {
	if l := len(p); l < MinPasswordLen {
		return fmt.Errorf("password too short, expecting at least %d bytes", MinPasswordLen)
	}

	return nil
}

func randUuid() string {
	return uuid.New().String()
}
