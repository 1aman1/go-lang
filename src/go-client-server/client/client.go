package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	createItem("New Item")
	getItems()
	getItem(1)
	updateItem(1, "Updated Item")
	deleteItem(1)
}

func createItem(name string) {
	item := Item{Name: name}
	data, _ := json.Marshal(item)
	resp, err := http.Post("http://localhost:8080/items", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Failed to create item: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Create response:", string(body))
}

func getItems() {
	resp, err := http.Get("http://localhost:8080/items")
	if err != nil {
		log.Fatalf("Failed to get items: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Get items response:", string(body))
}

func getItem(id int) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/items/%d", id))
	if err != nil {
		log.Fatalf("Failed to get item: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Get item response:", string(body))
}

func updateItem(id int, name string) {
	item := Item{Name: name}
	data, _ := json.Marshal(item)
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/items/%d", id), bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to update item: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Update response:", string(body))
}

func deleteItem(id int) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/items/%d", id), nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to delete item: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Delete response: No Content")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Delete response:", string(body))
	}
}
