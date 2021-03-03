package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handle preflight
func preflight(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Read all existing todos
func readTodos(w http.ResponseWriter, r *http.Request) {
	todos := Todos().GetAll()
	json.NewEncoder(w).Encode(todos)
}

// Create a new todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Todos().Create(todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// Read an existing todo
func readTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todo := Todos().Get(vars["id"])

	if todo != nil {
		json.NewEncoder(w).Encode(todo)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// Update an existing todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Delete an existing todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	success := Todos().Delete(vars["id"])

	if success == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete all existing todos
func deleteTodos(w http.ResponseWriter, r *http.Request) {
	success := Todos().DeleteAll()

	if success == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func configureRoutes(router *mux.Router) {
	// Preflight routes
	router.HandleFunc("/todos", preflight).Methods(http.MethodOptions)
	router.HandleFunc("/todos/{id}", preflight).Methods(http.MethodOptions)

	// Todo routes
	router.HandleFunc("/todos", readTodos).Methods(http.MethodGet)
	router.HandleFunc("/todos", createTodo).Methods(http.MethodPost)
	router.HandleFunc("/todos", deleteTodos).Methods(http.MethodDelete)
	router.HandleFunc("/todos/{id}", readTodo).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", updateTodo).Methods(http.MethodPatch)
	router.HandleFunc("/todos/{id}", deleteTodo).Methods(http.MethodDelete)
}
