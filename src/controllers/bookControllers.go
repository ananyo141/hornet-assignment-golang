package controllers

import (
	"backend/src/models"
	"backend/src/utils"
	"log"
	"strings"

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

func AddBook(ctx *fiber.Ctx) error {
	book := new(models.Book)
	if err := ctx.BodyParser(book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, err.Error(), nil))
	}

	errs := utils.Validate.Struct(book)
	if errs != nil {
		validationErrors := utils.NormalizeErrors(errs)
		// Return if there are validation errors
		if len(validationErrors) > 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, "Unprocessable Entity", validationErrors))
		}
	}

	// Determine the correct file path based on user role
	filePath := utils.UserBooksFilePath

	// Check if the book already exists
	exists, err := utils.BookExists(filePath, book.Name)
	if err != nil {
    log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse(false, "Failed to check if book exists.", nil))
	}
	if exists {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, "Book already exists.", nil))
	}

	// Add the book to the CSV file
	err = utils.AddBookToCSV(filePath, *book)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse(false, "Failed to add the book.", err))
	}

	return ctx.JSON(utils.HttpResponse(true, "Book added successfully.", nil))
}

func DeleteBook(ctx *fiber.Ctx) error {
	bookName := ctx.Params("name") // Assuming the book name is passed as a URL parameter
	if bookName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, "Book name is required.", nil))
	}
	bookName = strings.Replace(bookName, "%20", " ", -1) // Replace %20 with space (if any)

	exists, err := utils.BookExists(utils.UserBooksFilePath, bookName)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse(false, "Failed to check if book exists.", nil))
	}
	if !exists {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, "Book does not exist", nil))
	}

	// Determine the correct file path based on user role
	// Delete the book from the CSV file
	delerr := utils.DeleteBook(utils.UserBooksFilePath, bookName)
	if delerr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse(false, "Failed to delete the book.", nil))
	}

	return ctx.JSON(utils.HttpResponse(true, "Book deleted successfully.", nil))
}
