package db

import (
	"appGin/internal/models"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib" // Use pgx driver for PostgreSQL
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	// Fetch the database connection string from environment variables
	dsn := os.Getenv("DATABASE_URL") // Make sure the DATABASE_URL environment variable is set
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// AutoMigrate the User model
	if err := db.AutoMigrate(&models.Users{}); err != nil {
		log.Fatalf("Failed to auto-migrate the User model: %v", err)
	}

	log.Println("Database migrated successfully!")
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal("Error closing the database connection:", err)
	}
}
