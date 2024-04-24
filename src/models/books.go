package models

import (
	"time"
	"github.com/go-playground/validator/v10"
)

// Struct to hold book information
type Book struct {
	Name            string `json:"name" validate:"required"`
	Author          string `json:"author" validate:"required"`
	PublicationYear string `json:"publication_year" validate:"required,year"`
}

// Custom validation function for the publication year.
func YearValidation(fl validator.FieldLevel) bool {
	year := fl.Field().String()
	layout := "2006" // Go's reference time format
	parsedYear, err := time.Parse(layout, year)
	if err != nil {
		return false
	}
	currentYear := time.Now().Year()
	if parsedYear.Year() >= 1000 && parsedYear.Year() <= currentYear {
		return true
	}
	return false
}

