package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib" // Use pgx driver for PostgreSQL
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")

	// Initialize the database connection
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing the database connection:", err)
		}
	}(db)

	// Check if the connection is established
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to DB:", err)
	} else {
		fmt.Println("Connected to the database successfully! 🎉")
	}
}
