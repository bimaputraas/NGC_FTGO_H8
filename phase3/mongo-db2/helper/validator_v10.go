package helper

import "github.com/go-playground/validator/v10"

func ValidateStruct(myStruct interface{}) error{
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(myStruct)
	if err != nil {
		return err
	}
	return nil
}