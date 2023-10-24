package common

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(input interface{}) error {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}
	return nil
}
