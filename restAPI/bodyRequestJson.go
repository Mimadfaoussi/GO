package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a struct to match the JSON payload
type ServiceStatus struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

func statusUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var status ServiceStatus

	// Parse the JSON body into the struct
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Close the body when done

	// Respond with the received data
	fmt.Fprintf(w, "Service: %s, Status: %s\n", status.Service, status.Status)
}

func main() {
	http.HandleFunc("/update-status", statusUpdateHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
