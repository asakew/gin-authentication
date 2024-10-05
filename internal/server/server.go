package server

import (
	"appGin/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	gin.SetMode(gin.ReleaseMode) // Set gin to release mode

	// Set up Gin router
	//r := gin.New() // Creates a new Gin engine without default middleware
	r := gin.Default()

	r.Use(gin.Logger())   // Use Logger middleware
	r.Use(gin.Recovery()) // Use Recovery middleware

	// _________________________________________________________________

	// _________________________________________________________________

	// Register routes
	routes.HTMLRendering(r)
	//routes.SetupAuth(r) + verifyRecaptcha.go

	// Start the server
	log.Println("Server running: http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run the server:", err)
	}
}
