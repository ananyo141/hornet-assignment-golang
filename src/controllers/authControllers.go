package controllers

import (
	"backend/src/models"
	"backend/src/utils"

	"github.com/gofiber/fiber/v2"
)


func Login(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, err.Error(), nil))
	}

	errs := utils.Validate.Struct(user)
	if errs != nil {
		validationErrors := utils.NormalizeErrors(errs)
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
