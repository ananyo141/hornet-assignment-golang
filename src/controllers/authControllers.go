package controllers

import (
	"backend/src/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type User struct {
	Name    string `validate:"required"`
	Email   string `validate:"required,email"`
	IsAdmin string `validate:"omitempty,oneof=true false"`
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func Login(ctx *fiber.Ctx) error {
	user := new(User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, err.Error(), nil))
	}

	validationErrors := []ErrorResponse{}
	errs := validate.Struct(user)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true
			validationErrors = append(validationErrors, elem)
		}

		// Return if there are validation errors
		if len(validationErrors) > 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, "Unprocessable Entity", validationErrors))
		}
	}

	var userRole string
	if user.IsAdmin == "true" {
		userRole = "admin"
	} else {
		userRole = "user"
	}

	token, err := utils.GenerateJWT(user.Name, userRole)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse(false, err.Error(), nil))
	}

	// Correctly return the generated JWT token
	return ctx.JSON(utils.HttpResponse(true, "Logged in successfully", fiber.Map{
		"token":    token,
		"userType": userRole,
	}))
}
