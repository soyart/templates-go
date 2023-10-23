package todo

import (
	"fmt"
	"reflect"
	"testing"

	"example.com/servicehex/domain/core"
	"example.com/servicehex/domain/datagateway/mock"
	"example.com/servicehex/domain/entity"
)

func mockTodo(i int, baseTodoID string, userID string) entity.Todo {
	return entity.Todo{
		ID:     fmt.Sprintf("%s-%d", baseTodoID, i),
		UserID: userID,
		Text:   fmt.Sprintf("foobar-%d", i),
	}
}

func TestTodoCrud(t *testing.T) {
	baseTodoID, userID := "fooTodo", "user1"
	repo := mock.NewMockTodoRepo(nil)
	todo := mockTodo(0, baseTodoID, userID)

	if err := createTodo(nil, repo, todo); err != nil {
		t.Logf("Unexpected error from createTodo")
		t.Errorf("Error: %s", err.Error())
	}

	savedTodo, err := getTodoByID(nil, repo, userID, todo.ID)
	if err != nil {
		t.Logf("Unexpected error from getTodoByID")
		t.Errorf("Error: %s", err.Error())
	}

	if !reflect.DeepEqual(todo, savedTodo) {
		t.Logf("Unexpected value")
		t.Logf("Expected: %+v", todo)
		t.Logf("Actual: %+v", savedTodo)
	}

	numTodos := 10
	for i := 1; i < numTodos; i++ {
		err := createTodo(nil, repo, mockTodo(i, baseTodoID, userID))
		if err != nil {
			t.Logf("Unexpected error when looping to create mock todo")
			t.Errorf("Error: %s", err.Error())
		}
	}

	todos, err := getTodos(nil, repo, userID, core.SortNoSort)
	if l := len(todos); l < numTodos {
		t.Logf("Unexpected number of mock todos")
		t.Logf("Expected: %d", numTodos)
		t.Logf("Actual: %d", l)
		t.Error()
	}

	for i := range todos {
		expected := mockTodo(i, baseTodoID, userID)
		actual := todos[i]

		if !reflect.DeepEqual(expected, actual) {
			t.Logf("Unexpected value")
			t.Logf("Expected: %+v", expected)
			t.Logf("Actual: %+v", actual)
		}

		deleteTodoByID(nil, repo, userID, actual.ID)
	}

	todos, err = getTodos(nil, repo, userID, core.SortNoSort)
	if err != nil {
		t.Logf("Unexpected error after delete")
		t.Errorf("Error: %s", err.Error())
	}
	if l := len(todos); l != 0 {
		t.Logf("Unexpected todo length")
		t.Errorf("Expecting %d, got %d", 0, l)
	}

	for i := 0; i < numTodos; i++ {
		todoID := mockTodo(i, baseTodoID, userID).ID

		todo, err := getTodoByID(nil, repo, userID, todoID)
		if err != nil {
			t.Logf("Found error when get after delete")
			t.Logf("Error: %s", err.Error())
		}

		if todo.ID == todoID {
			t.Errorf("Unexpected value (todoID) after delete")
		}

		if todo.Text != "" {
			t.Errorf("Unexpected value (text) after delete")
		}
	}
}
