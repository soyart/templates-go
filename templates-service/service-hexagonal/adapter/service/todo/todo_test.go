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
	}
}
