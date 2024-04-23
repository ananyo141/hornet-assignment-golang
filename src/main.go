package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"backend/src/middlewares"
	"backend/src/routes"
	"backend/src/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})
	app.Use(recover.New())

	// Routes
	routes.AuthRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(utils.HttpResponse(true, "Server is running.", fiber.Map{}))
	})

	app.Listen(":3000")
}
