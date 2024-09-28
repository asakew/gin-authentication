package server

import (
	"appGin/internal/db"
	"appGin/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func Run() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db.InitDB()
	defer db.CloseDB()

	gin.SetMode(gin.ReleaseMode) // Set gin to release mode

	// Set up Gin router
	r := gin.New() // Creates a new Gin engine without default middleware

	// Set trusted proxies (replace with actual IPs of your trusted proxies)
	err = r.SetTrustedProxies([]string{"10.0.0.0/8", "192.168.1.1"})
	if err != nil {
		log.Fatal("Error setting trusted proxies:", err)
	}

	r.Use(gin.Logger())   // Use Logger middleware
	r.Use(gin.Recovery()) // Use Recovery middleware

	// Register routes
	routes.TemplatesRoutes(r)

	// Start the server
	log.Println("Server running: http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run the server:", err)
	}
}
