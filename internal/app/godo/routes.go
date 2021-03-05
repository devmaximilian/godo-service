package godo

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/devmaximilian/godo-service/internal/pkg/todo"
)

// Handle preflight
func preflight(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Read all existing todos
func readTodos(w http.ResponseWriter, r *http.Request) {
	todos := todo.Todos().GetAll()
	json.NewEncoder(w).Encode(todos)
}

// Create a new todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	var value todo.Todo
	err := json.NewDecoder(r.Body).Decode(&value)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo.Todos().Create(value)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(value)
}

// Read an existing todo
func readTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todo := todo.Todos().Get(vars["id"])

	if todo != nil {
		json.NewEncoder(w).Encode(todo)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// Update an existing todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := todo.Todos().Get(vars["id"])

	var newValue todo.Todo
	err := json.NewDecoder(r.Body).Decode(&newValue)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo.Todos().Update(value, &newValue)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(value)
}

// Delete an existing todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	success := todo.Todos().Delete(vars["id"])

	if success == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete all existing todos
func deleteTodos(w http.ResponseWriter, r *http.Request) {
	success := todo.Todos().DeleteAll()

	if success == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
