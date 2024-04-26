package db

import (
	"backend/src/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

// ConnectDB connects to the database and returns the connection
func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to the database: ", err)
		return nil, err
	}
	log.Println("Connected to the database")
	return db, nil
}

// MigrateDB migrates the database schema
func MigrateDB(db *gorm.DB) error {
	log.Println("Database migrated")

	// Add a custom, case-insensitive unique constraint on the 'name' column
	err := db.AutoMigrate(&models.Book{})
	db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS book_name_unique_idx ON books (LOWER(name));")
	return err
}
