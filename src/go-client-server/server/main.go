package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	items   = []Item{}
	nextID  = 1
	itemsMu sync.Mutex
)

func main() {
	http.HandleFunc("/items", handleItems)
	http.HandleFunc("/items/", handleItem)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		createItem(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getItem(w, r, id)
	case http.MethodPut:
		updateItem(w, r, id)
	case http.MethodDelete:
		deleteItem(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getItems(w http.ResponseWriter, r *http.Request) {
	itemsMu.Lock()
	defer itemsMu.Unlock()
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	itemsMu.Lock()
	defer itemsMu.Unlock()
	item.ID = nextID
	nextID++
	items = append(items, item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func getItem(w http.ResponseWriter, r *http.Request, id int) {
	itemsMu.Lock()
	defer itemsMu.Unlock()
	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var updatedItem Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	itemsMu.Lock()
	defer itemsMu.Unlock()
	for i, item := range items {
		if item.ID == id {
			items[i].Name = updatedItem.Name
			json.NewEncoder(w).Encode(items[i])
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request, id int) {
	itemsMu.Lock()
	defer itemsMu.Unlock()
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}
