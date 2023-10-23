package todo

import (
	"fmt"
	"reflect"
	"testing"

	"example.com/servicehex/domain/core"
	"example.com/servicehex/domain/datagateway/mock"
	"example.com/servicehex/domain/entity"
)

func mockTodo(i int, baseTodoId string, userId string) entity.Todo {
	return entity.Todo{
		Id:     fmt.Sprintf("%s-%d", baseTodoId, i),
		UserId: userId,
		Text:   fmt.Sprintf("foobar-%d", i),
	}
}

func TestTodoCrud(t *testing.T) {
	baseTodoId, userId := "fooTodo", "user1"
	repo := mock.NewMockTodoRepo(nil)
	todo := mockTodo(0, baseTodoId, userId)

	if err := createTodo(nil, repo, todo); err != nil {
		t.Logf("Unexpected error from createTodo")
		t.Errorf("Error: %s", err.Error())
	}

	savedTodo, err := getTodoById(nil, repo, userId, todo.Id)
	if err != nil {
		t.Logf("Unexpected error from getTodoById")
		t.Errorf("Error: %s", err.Error())
	}

	if !reflect.DeepEqual(todo, savedTodo) {
		t.Logf("Unexpected value")
		t.Logf("Expected: %+v", todo)
		t.Logf("Actual: %+v", savedTodo)
	}

	numTodos := 10
	for i := 1; i < numTodos; i++ {
		err := createTodo(nil, repo, mockTodo(i, baseTodoId, userId))
		if err != nil {
			t.Logf("Unexpected error when looping to create mock todo")
			t.Errorf("Error: %s", err.Error())
		}
	}

	todos, err := getTodos(nil, repo, userId, core.SortNoSort)
	if l := len(todos); l < numTodos {
		t.Logf("Unexpected number of mock todos")
		t.Logf("Expected: %d", numTodos)
		t.Logf("Actual: %d", l)
		t.Error()
	}

	for i := range todos {
		expected := mockTodo(i, baseTodoId, userId)
		actual := todos[i]

		if !reflect.DeepEqual(expected, actual) {
			t.Logf("Unexpected value")
			t.Logf("Expected: %+v", expected)
			t.Logf("Actual: %+v", actual)
		}

		deleteTodoById(nil, repo, userId, actual.Id)
	}

	todos, err = getTodos(nil, repo, userId, core.SortNoSort)
	if err != nil {
		t.Logf("Unexpected error after delete")
		t.Errorf("Error: %s", err.Error())
	}
	if l := len(todos); l != 0 {
		t.Logf("Unexpected todo length")
		t.Errorf("Expecting %d, got %d", 0, l)
	}

	for i := 0; i < numTodos; i++ {
		todoId := mockTodo(i, baseTodoId, userId).Id

		todo, err := getTodoById(nil, repo, userId, todoId)
		if err != nil {
			t.Logf("Found error when get after delete")
			t.Logf("Error: %s", err.Error())
		}

		if todo.Id == todoId {
			t.Errorf("Unexpected value (todoId) after delete")
		}

		if todo.Text != "" {
			t.Errorf("Unexpected value (text) after delete")
		}
	}
}
