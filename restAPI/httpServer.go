package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Data struct {
	Message string `json:"message"`
	Status int `json:"status"`
}


func homeHandle(w http.ResponseWriter, r *http.Request) {
	data := Data {
		Message : "This is the life",
		Status : 200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}


func main() {
	http.HandleFunc("/home", homeHandle)
	fmt.Println("running httpServer on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
