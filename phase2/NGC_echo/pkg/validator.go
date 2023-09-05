package pkg

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(data)
	if err != nil {
		return err
	}

	return nil
}
