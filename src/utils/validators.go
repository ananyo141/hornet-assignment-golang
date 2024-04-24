package utils

import (
	"backend/src/models"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// Register custom validation functions
// Add other validation functions here
func init() {
	Validate.RegisterValidation("year", models.YearValidation)
}
