package helper

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) error {
	validator := validator.New()
	err := validator.Struct(data)
	return err
}