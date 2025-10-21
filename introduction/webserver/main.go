package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from a simple listener!"))
	})

	// Get port from environment variable, default to 8082
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	addr := ":" + port
	log.Printf("Starting simple webserver on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}