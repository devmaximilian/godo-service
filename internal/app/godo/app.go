package godo

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/devmaximilian/godo-service/internal/pkg/middleware"
)

type application struct {
	router *mux.Router
}

// Create a new application instance
func NewApp() *application {
	a := application{
		router: mux.NewRouter(),
	}
	a.configure()
	return &a
}

// Configure routes and middlewares
func (a *application) configure() {
	// Register middlewares
	a.router.Use(middleware.CorsMiddleware)
	a.router.Use(middleware.LoggingMiddleware)

	// Register routes

	// Preflight
	a.router.HandleFunc("/todos", preflight).Methods(http.MethodOptions)
	a.router.HandleFunc("/todos/{id}", preflight).Methods(http.MethodOptions)

	// Todo
	a.router.HandleFunc("/todos", readTodos).Methods(http.MethodGet)
	a.router.HandleFunc("/todos", createTodo).Methods(http.MethodPost)
	a.router.HandleFunc("/todos", deleteTodos).Methods(http.MethodDelete)
	a.router.HandleFunc("/todos/{id}", readTodo).Methods(http.MethodGet)
	a.router.HandleFunc("/todos/{id}", updateTodo).Methods(http.MethodPatch)
	a.router.HandleFunc("/todos/{id}", deleteTodo).Methods(http.MethodDelete)
}

// Run application
func (a *application) Run() {
	log.Fatal(http.ListenAndServe(":8080", a.router))
}
