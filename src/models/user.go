package models

type User struct {
	Name    string `validate:"required"`
	Email   string `validate:"required,email"`
	IsAdmin string `validate:"omitempty,oneof=true false"`
}

