package routes

import (
	"backend/src/controllers"
	"backend/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	bookRoutes := app.Group("/books")
	bookRoutes.Get("/", middlewares.AuthRequired(), controllers.GetBooks)
	// bookRoutes.Post("/", controllers.CreateBook)
	// bookRoutes.Get("/:id", controllers.GetBook)
	// bookRoutes.Put("/:id", controllers.UpdateBook)
	// bookRoutes.Delete("/:id", controllers.DeleteBook)

}
