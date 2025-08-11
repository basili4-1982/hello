package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response структура ответа
type Response struct {
	Message string `json:"message"`
}

// handler — простой обработчик
func handler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Hello, CI/CD with Go!"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
