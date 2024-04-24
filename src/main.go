package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"backend/src/middlewares"
	"backend/src/routes"
	"backend/src/utils"
)

func init() {
	// populate the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Println("No .env file loaded, reading from system environment variables, recheck all the environment variables are set.")
	}
	utils.JwtSecret = []byte(os.Getenv("JWT_KEY"))
	utils.AdminBooksFilePath = os.Getenv("ADMIN_FILE")
	utils.UserBooksFilePath = os.Getenv("REGULAR_FILE")
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})
	app.Use(recover.New())

	// Routes
	routes.AuthRoutes(app)
	routes.BookRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(utils.HttpResponse(true, "Server is running.", fiber.Map{}))
	})

	app.Listen(":3000")
}
