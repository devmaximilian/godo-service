package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Register CORS middleware
	router.Use(corsMiddleware)

	// Register routes
	configureRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
