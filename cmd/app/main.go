package main

import (
	"appGin/internal/config"
	"appGin/internal/repo"
	"appGin/internal/server"
)

func main() {
	config.InitDB()
	config.InitRedis()
	repo.DB = config.DB

	// Start the server
	server.Run()
}
