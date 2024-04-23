package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"backend/src/utils"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	// Send custom error page
	err = ctx.Status(code).JSON(utils.HttpResponse(false, "Error", fiber.Map{}))
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}
