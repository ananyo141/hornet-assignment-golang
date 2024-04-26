package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"log"
	"os"

	"backend/src/db"
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

	_db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	db.Db = _db
	if err := db.MigrateDB(db.Db); err != nil {
		log.Fatal(err)
	}
	log.Println("Database initialized")
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
