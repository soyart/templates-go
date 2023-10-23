package core

import "fmt"

type ErrWrongUserID struct {
	todoID string
	userID string
}

func (e ErrWrongUserID) Error() string {
	return fmt.Sprintf("no such todoID %s with userID %s", e.todoID, e.userID)
}

func WrongUserID(todoID, userID string) error {
	return ErrWrongUserID{
		todoID: todoID,
		userID: userID,
	}
}
