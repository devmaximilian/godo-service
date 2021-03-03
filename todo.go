package main

import (
	"strconv"
)

type Todo struct {
	Id        string `json:"-"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Text      string `json:"text"`
}

// In-memory storage type
type todos struct {
	items  map[string]*Todo
	nextId uint64
}

// Singleton instance
var (
	t *todos
)

func Todos() *todos {
	if t == nil {
		t = &todos{
			items:  make(map[string]*Todo),
			nextId: 1,
		}
	}
	return t
}

// – CRUD Methods

// Get todo
func (t *todos) Get(id string) *Todo {
	for key, item := range t.items {
		if key == item.Id {
			return item
		}
	}
	return nil
}

// Get all todos
func (t *todos) GetAll() []*Todo {
	items := []*Todo{}
	for _, item := range t.items {
		items = append(items, item)
	}
	return items
}

// Get todo
func (t *todos) Delete(id string) bool {
	for key, item := range t.items {
		if key == item.Id {
			delete(t.items, key)
			return true
		}
	}
	return false
}

// Delete all todos
func (t *todos) DeleteAll() bool {
	for key := range t.items {
		delete(t.items, key)
	}
	return true
}

// Create todo
func (t *todos) Create(todo Todo) *Todo {
	todo.Id = strconv.FormatUint(t.nextId, 16)
	todo.Url = "/todos/" + todo.Id
	t.items[todo.Id] = &todo
	t.nextId++
	return &todo
}

// Update todo
func (t *todos) Update(todo *Todo, updated *Todo) *Todo {
	todo.Title = updated.Title
	todo.Order = updated.Order
	todo.Text = updated.Text
	todo.Completed = updated.Completed
	return todo
}
