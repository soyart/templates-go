package common

import (
	"errors"
	"fmt"
)

var (
	ErrCreateTodo = errors.New("failed to create todo")
)

func WrapError(errContext string, err error) error {
	if err == nil {
		return err
	}
	return fmt.Errorf("%s: %w", errContext, err)
}
