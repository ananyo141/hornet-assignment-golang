package controllers

import (
	"backend/src/db"
	"backend/src/models"
	"backend/src/utils"
	"errors"

	"gorm.io/gorm"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetBooks(ctx *fiber.Ctx) error {
	var books []models.Book
	result := db.Db.Find(&books)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse(false, "Error retrieving books.", nil))
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

	// Check if the book already exists
	err := db.Db.Create(&book).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.Println(pgErr.Message)
			log.Println(pgErr.Code)
			if pgErr.Code == "23505" {
				return ctx.Status(fiber.StatusConflict).JSON(utils.HttpResponse(false, "Book already exists.", nil))
			} else {
				return ctx.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse(false, err.Error(), nil))
			}
		}
	}
	return ctx.JSON(utils.HttpResponse(true, "Book added successfully.", book))
}

func UpdateBook(ctx *fiber.Ctx) error {
	bookID := ctx.Params("id")
	if bookID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, "Book ID is required.", nil))
	}

	// Parse the updated book data
	var updateData models.Book
	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, err.Error(), nil))
	}

	// Find the book by ID and update it
	var book models.Book
	result := db.Db.First(&book, bookID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(utils.HttpResponse(false, "Book not found.", nil))
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse(false, "Error finding book.", nil))
	}

	// Perform the update
	db.Db.Model(&book).Updates(updateData)

	return ctx.JSON(utils.HttpResponse(true, "Book updated successfully.", book))
}

func DeleteBook(ctx *fiber.Ctx) error {
	bookID := ctx.Params("id")
	if bookID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse(false, "Book ID is required.", nil))
	}

	// Attempt to delete the book by ID
	result := db.Db.Delete(&models.Book{}, bookID)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse(false, "Error deleting book.", nil))
	}

	// Check if any rows were affected (i.e., if the book existed)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(utils.HttpResponse(false, "Book not found.", nil))
	}

	return ctx.JSON(utils.HttpResponse(true, "Book deleted successfully.", nil))
}
