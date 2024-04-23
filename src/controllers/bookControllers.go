package controllers

import (
	"backend/src/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(ctx *fiber.Ctx) error {
	books, error := utils.LoadBooksFromCSV(utils.UserBooksFilePath)
	user := ctx.Locals("user").(*utils.Claims)
	if error != nil {
		log.Println(error)
		return ctx.JSON(utils.HttpResponse(false, "Error loading books.", fiber.Map{}))
	}
	if user.UserType == "admin" {
		adminBooks, error2 := utils.LoadBooksFromCSV(utils.AdminBooksFilePath)
		if error2 != nil {
			log.Println(error2)
			return ctx.JSON(utils.HttpResponse(false, "Error loading admin books.", fiber.Map{}))
		}
		books = append(books, adminBooks...)
	}
	return ctx.JSON(utils.HttpResponse(true, "Books retrieved successfully.", books))
}
