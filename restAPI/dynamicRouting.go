package main

import (
	"fmt"
	"net/http"
	"strings"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the service name from the URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[1] != "status" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	service := pathParts[2] // Get the service name

	// Simulate returning the status
	fmt.Fprintf(w, "Status of %s: OK\n", service)
}

func main() {
	http.HandleFunc("/status/", statusHandler) // Note the trailing slash
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
