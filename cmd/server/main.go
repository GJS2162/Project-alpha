package main

import (
	"log"
	"net/http"

	"logistics-platform/internal/api"
	"logistics-platform/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	if err := database.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Set up Gin router
	router := gin.Default()

	// Set up routes
	api.SetupRoutes(router)

	// Start the server
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
