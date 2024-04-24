package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"time"

	"github.com/go-playground/validator/v10"
)

var AdminBooksFilePath = "src/data/adminUser.csv"
var UserBooksFilePath = "src/data/regularUser.csv"

// Define a struct to hold book information
type Book struct {
	Name            string `json:"name" validate:"required"`
	Author          string `json:"author" validate:"required"`
	PublicationYear string `json:"publication_year" validate:"required,year"`
}

// Custom validation function for the publication year.
func YearValidation(fl validator.FieldLevel) bool {
	year := fl.Field().String()
	layout := "2006" // Go's reference time format
	parsedYear, err := time.Parse(layout, year)
	if err != nil {
		return false
	}
	currentYear := time.Now().Year()
	if parsedYear.Year() >= 1000 && parsedYear.Year() <= currentYear {
		return true
	}
	return false
}

// Function to load books from the CSV file
func LoadBooksFromCSV(filePath string) ([]Book, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ',' // Set the delimiter
	reader.TrimLeadingSpace = true

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var books []Book
	for i, line := range lines {
		if i == 0 { // Skip the header line
			continue
		}
		if len(line) >= 3 {
			books = append(books, Book{
				Name:            line[0],
				Author:          line[1],
				PublicationYear: line[2],
			})
		}
	}

	return books, nil
}

func AddBookToCSV(filePath string, book Book) error {
	// Open the file in read-write mode, and create it if it does not exist.
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the new book entry.
	err = writer.Write([]string{book.Name, book.Author, book.PublicationYear})
	if err != nil {
		return err
	}

	return nil
}

func BookExists(filePath string, newBook string) (bool, error) {
	// Open the CSV file in read-only mode.
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Create a CSV reader.
	reader := csv.NewReader(file)

	// Read all records.
	records, err := reader.ReadAll()
	if err != nil {
		return false, err
	}

	// Loop through the records to find if the book already exists.
	for _, record := range records {
		if len(record) >= 3 {
			name := record[0]
			if strings.ToLower(name) == strings.ToLower(newBook) {
				return true, nil // Book found.
			}
		}
	}

	return false, nil // Book not found.
}

func DeleteBook(filePath, bookName string) error {
	// Open the original CSV file in read-only mode.
	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create a temporary file to write the filtered records.
	tmpFilePath := filePath + ".tmp"
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	reader := csv.NewReader(inputFile)
	writer := csv.NewWriter(tmpFile)

	// Flag to check if a book has been found and deleted.
	bookDeleted := false

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// If the book name matches, skip writing this record (delete it).
		if strings.ToLower(record[0]) == strings.ToLower(bookName) {
			bookDeleted = true
			continue
		}

		// Write the record to the temporary file.
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	writer.Flush()

	if !bookDeleted {
		log.Printf("Book \"%s\" not found. No deletion performed.\n", bookName)
		// Remove the temporary file since no changes were made.
		os.Remove(tmpFilePath)
		return nil
	}

	// Close files before renaming to ensure all data is flushed and files are not locked.
	inputFile.Close()
	tmpFile.Close()

	// Replace the original file with the temporary file.
	if err := os.Rename(tmpFilePath, filePath); err != nil {
		return err
	}

	log.Printf("Book \"%s\" deleted successfully.\n", bookName)
	return nil
}
