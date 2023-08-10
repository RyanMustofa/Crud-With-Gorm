package helpers

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field string `json:"field"`
	Tag string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func Validation[T any](payload T) []ErrorResponse {
	validate := validator.New()
	var error []ErrorResponse

	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Error()
			error = append(error,element)
		}
	}

	return error
}
