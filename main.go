package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := Response{Message: "OK", Status: http.StatusOK}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := Response{Message: "Hello, World!", Status: http.StatusOK}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/hello", helloHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}