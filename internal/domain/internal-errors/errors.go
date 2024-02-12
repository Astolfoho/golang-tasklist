package internalerrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var ErrInternalServerError = errors.New("internal server error")

func ValidateStruct(obj interface{}) validator.FieldError {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	error := err.(validator.ValidationErrors)[0]
	return error
}
