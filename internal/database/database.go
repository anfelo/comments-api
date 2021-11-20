package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
	)

	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
