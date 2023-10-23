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
	GetTodoByID(ctx context.Context, userID, todoID string) (entity.Todo, error)
	GetTodos(ctx context.Context, userID string, sortMode EnumSortMode) ([]entity.Todo, error)
	GetExpiredTodos(ctx context.Context, userID string, cutoff time.Time, sortMode EnumSortMode) ([]entity.Todo, error)
	MatchTodoTextPattern(ctx context.Context, userID, pattern string, sortMode EnumSortMode) ([]entity.Todo, error)
	ExpireTodoByID(ctx context.Context, userID, todoID string) error
	DeleteTodoByID(ctx context.Context, userID, todoID string) error
	DeleteTodos(ctx context.Context, userID string) error
}

type PortUser interface {
	Register(ctx context.Context, user entity.User) error
	Login(ctx context.Context, username string, password []byte) (entity.User, error)
	ChangePassword(ctx context.Context, userID string, password []byte, newPassword []byte) error
	DeleteUser(ctx context.Context, userID string) error
}

func ExpireTodo(todo *entity.Todo) {
	todo.Expired, todo.ExpiredAt = true, time.Now()
}

func TodoID() string {
	return randUuid()
}

func UserID(username string) string {
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
