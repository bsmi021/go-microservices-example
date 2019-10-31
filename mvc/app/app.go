package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
	
)

func init() {
	router = gin.Default() // returns engine with logger and recovery middleware already attached
	// router = gin.New() // returns a blank instance without middleware
}

// StartApp Starts the application which is handling web calls
func StartApp() {
	mapUrls()	
	

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
