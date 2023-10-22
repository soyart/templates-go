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

func pointToExactTodo(todos []entity.Todo, userId, todoId string) *entity.Todo {
	for i := range todos {
		t := &todos[i]
		if t.Id == todoId && t.UserId == userId {
			return t
		}
	}

	return nil
}

func (m *mockRepoTodo) CreateTodo(ctx context.Context, todo entity.Todo) error {
	m.todos = append(m.todos, todo)
	return nil
}

func (m *mockRepoTodo) GetTodo(ctx context.Context, userId string, todoId string) (entity.Todo, error) {
	for i := range m.todos {
		todo := &m.todos[i]

		if todo.Id == todoId && todo.UserId == userId {
			return *todo, nil
		}
	}

	return entity.Todo{}, nil
}

func (m *mockRepoTodo) GetTodos(ctx context.Context, userId string) ([]entity.Todo, error) {
	return filterUserTodos(m.todos, userId)
}

func (m *mockRepoTodo) UpdateTodo(ctx context.Context, userId string, todoId string, update entity.Todo) error {
	target := pointToExactTodo(m.todos, userId, todoId)
	if target == nil {
		return core.WrongUserId(todoId, userId)
	}

	*target = update
	return nil
}

func (m *mockRepoTodo) DeleteTodo(ctx context.Context, userId string, todoId string) error {
	for i := range m.todos {
		t := &m.todos[i]
		if t.Id == todoId && t.UserId == userId {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)

			return nil
		}
	}

	return fmt.Errorf("no such todo %s for user %s", todoId, userId)
}

func (m *mockRepoTodo) DeleteTodos(ctx context.Context, userId string) error {
	var retained []entity.Todo
	for i := range m.todos {
		t := &m.todos[i]

		if t.UserId != userId {
			retained = append(retained, *t)
		}
	}

	m.todos = retained
	return nil
}

func filterUserTodos(todos []entity.Todo, userId string) ([]entity.Todo, error) {
	if todos == nil {
		return nil, fmt.Errorf("got nil todos")
	}

	var userTodos []entity.Todo
	for i := range todos {
		t := todos[i]

		if t.UserId == userId {
			userTodos = append(userTodos, t)
		}
	}

	return userTodos, nil
}
