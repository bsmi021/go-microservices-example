package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default() // returns engine with logger and recovery middleware already attached
	// router = gin.New() // returns a blank instance without middleware
}

// StartApp starts the MVC application which is handling web calls
func StartApp() {
	mapUrls()

	// Get port from environment variable, default to 8081
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	addr := ":" + port
	log.Printf("Starting MVC server on %s", addr)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
