package routes

import (
	auth "backend/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	// AuthRoutes handles all the authentication routes
	authRoutes := app.Group("/auth")
	authRoutes.Post("/login", auth.Login)
}
