package mock

import (
	"context"
	"fmt"

	"example.com/servicehex/domain/core"
	"example.com/servicehex/domain/datagateway"
	"example.com/servicehex/domain/entity"
)

type mockRepoTodo struct {
	todos []entity.Todo
}

func NewMockTodoRepo(initTodos []entity.Todo) datagateway.DataGatewayTodo {
	return &mockRepoTodo{todos: initTodos}
}

func pointToExactTodo(todos []entity.Todo, userID, todoID string) *entity.Todo {
	for i := range todos {
		t := &todos[i]
		if t.ID == todoID && t.UserID == userID {
			return t
		}
	}

	return nil
}

func (m *mockRepoTodo) CreateTodo(ctx context.Context, todo entity.Todo) error {
	m.todos = append(m.todos, todo)
	return nil
}

func (m *mockRepoTodo) GetTodo(ctx context.Context, userID string, todoID string) (entity.Todo, error) {
	for i := range m.todos {
		todo := &m.todos[i]

		if todo.ID == todoID && todo.UserID == userID {
			return *todo, nil
		}
	}

	return entity.Todo{}, nil
}

func (m *mockRepoTodo) GetTodos(ctx context.Context, userID string) ([]entity.Todo, error) {
	return filterUserTodos(m.todos, userID)
}

func (m *mockRepoTodo) UpdateTodo(ctx context.Context, userID string, todoID string, update entity.Todo) error {
	target := pointToExactTodo(m.todos, userID, todoID)
	if target == nil {
		return core.WrongUserID(todoID, userID)
	}

	*target = update
	return nil
}

func (m *mockRepoTodo) DeleteTodo(ctx context.Context, userID string, todoID string) error {
	for i := range m.todos {
		t := &m.todos[i]
		if t.ID == todoID && t.UserID == userID {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)

			return nil
		}
	}

	return fmt.Errorf("no such todo %s for user %s", todoID, userID)
}

func (m *mockRepoTodo) DeleteTodos(ctx context.Context, userID string) error {
	var retained []entity.Todo
	for i := range m.todos {
		t := &m.todos[i]

		if t.UserID != userID {
			retained = append(retained, *t)
		}
	}

	m.todos = retained
	return nil
}

func filterUserTodos(todos []entity.Todo, userID string) ([]entity.Todo, error) {
	if todos == nil {
		return nil, fmt.Errorf("got nil todos")
	}

	var userTodos []entity.Todo
	for i := range todos {
		t := todos[i]

		if t.UserID == userID {
			userTodos = append(userTodos, t)
		}
	}

	return userTodos, nil
}
