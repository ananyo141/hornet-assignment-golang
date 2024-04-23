package main

import "github.com/gofiber/fiber/v2"

import "backend/src/utils"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(utils.HttpResponse(true, "Hello, World!", nil))
	})

	app.Listen(":3000")
}
