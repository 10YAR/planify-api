package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data any) string {
	var errors string
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors += err.StructNamespace() + " " + err.Tag() + " " + err.Param() + " "
		}
	}
	return errors
}
