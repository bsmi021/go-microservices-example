package app

import (
	"github.com/bsmi021/go-microservices-example/mvc/controllers"
)

// mapUrls defines all of the routes in the application
func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}