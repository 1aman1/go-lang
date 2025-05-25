package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := HelloResponse{
		Message: "Helloo, world!",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

}

func main(){
	http.HandleFunc("/hello", helloHandler)

	log.Println("Server is running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Fatal(err)
	}
}

