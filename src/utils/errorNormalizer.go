package utils

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

// Function to format the validation errors for easy consumption
func NormalizeErrors(errs error) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	for _, err := range errs.(validator.ValidationErrors) {
		var elem ErrorResponse
		elem.FailedField = err.Field()
		elem.Tag = err.Tag()
		elem.Value = err.Value()
		elem.Error = true
		validationErrors = append(validationErrors, elem)
	}

	return validationErrors
}
