package middlewares

import (
	"backend/src/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract the token from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.HttpResponse(
				false,
				"Unauthorized: no token provided.",
				fiber.Map{},
			))
		}

		// Typically, the Authorization header is in the format `Bearer <token>`,
		// so we need to split by space and get the second part
		tokenString := strings.Split(authHeader, " ")[1]

		// Parse the token
		claims, err := utils.ValidateJWT(tokenString)
		// Make sure the token's algorithm matches what you expect

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(
				utils.HttpResponse(false, "Invalid or expired JWT", fiber.Map{}),
			)
		}

		// The token is valid and we have the claims,
		c.Locals("user", claims)
		return c.Next()
	}
}
