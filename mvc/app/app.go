package app

import (
	"log"
	"golang-microservices/mvc/controllers"
	"net/http"
)

// StartApp Starts the application which is handling web calls
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln(err)
		panic(err)
	}
}