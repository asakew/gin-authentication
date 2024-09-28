package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib" // Use pgx driver for PostgreSQL
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")

	DB, err = sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Check if the connection is established
	if err := DB.Ping(); err != nil {
		log.Fatal("Cannot connect to DB:", err)
	} else {
		log.Println("Connected to the database successfully!")
	}
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal("Error closing the database connection:", err)
	}
}
