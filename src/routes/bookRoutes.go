package routes

import (
	"backend/src/controllers"
	"backend/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	bookRoutes := app.Group("/books")
	bookRoutes.Get("/", middlewares.AuthRequired(), controllers.GetBooks)
	bookRoutes.Post("/", middlewares.AuthRequired(), middlewares.AdminAccess(), controllers.AddBook)
	bookRoutes.Delete("/:name", middlewares.AuthRequired(), middlewares.AdminAccess(), controllers.DeleteBook)
}
