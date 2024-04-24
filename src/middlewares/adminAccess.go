package middlewares

import (
	"backend/src/utils"

	"github.com/gofiber/fiber/v2"
)

func AdminAccess() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Locals("user") == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.HttpResponse(
				false,
				"Unauthorized: no token provided.",
				fiber.Map{},
			))
		}
		if c.Locals("user").(*utils.Claims).UserType != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(utils.HttpResponse(
				false,
				"Unauthorized: admin access required.",
				fiber.Map{},
			))
		}
		return c.Next()
	}
}
