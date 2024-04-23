package utils

import (
	"encoding/csv"
	"os"
)

var AdminBooksFilePath = "src/data/adminUser.csv"
var UserBooksFilePath = "src/data/regularUser.csv"

// Define a struct to hold book information
type Book struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationYear string `json:"publication_year"`
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
