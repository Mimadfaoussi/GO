package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the service name from the URL
	vars := mux.Vars(r) // Extracts dynamic parameters
	service := vars["service"]

	// Simulate returning the status
	fmt.Fprintf(w, "Status of %s: OK\n", service)
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define a dynamic route
	r.HandleFunc("/status/{service}", statusHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
