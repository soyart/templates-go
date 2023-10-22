package todo

import (
	"fmt"

	"example.com/servicehex/domain/entity"
)

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
