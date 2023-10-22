package core

import "fmt"

type ErrWrongUserId struct {
	todoId string
	userId string
}

func (e ErrWrongUserId) Error() string {
	return fmt.Sprintf("no such todoId %s with userId %s", e.todoId, e.userId)
}

func WrongUserId(todoId, userId string) error {
	return ErrWrongUserId{
		todoId: todoId,
		userId: userId,
	}
}
