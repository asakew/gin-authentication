package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib" // Use pgx driver for PostgreSQL
	"github.com/joho/godotenv"
	"log"
	"net/http"
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
		fmt.Println("Connected to the database successfully!")
	}

	gin.SetMode(gin.ReleaseMode) // Set gin to release mode

	// Set up Gin router ---------------------------------------------------------
	r := gin.New() // Creates a new Gin engine without default middleware

	// Set trusted proxies (replace with actual IPs of your trusted proxies)
	err = r.SetTrustedProxies([]string{"10.0.0.0/8", "192.168.1.1"})
	if err != nil {
		return
	}

	r.Use(gin.Logger())   // Use Logger middleware
	r.Use(gin.Recovery()) // Use Recovery middleware

	// Serve static files
	r.Static("/css", "./web/assets/css")
	r.Static("/js", "./web/assets/js")

	// Render HTML template
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Start the server
	log.Println("Server running: http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run the server:", err)
	}
}
