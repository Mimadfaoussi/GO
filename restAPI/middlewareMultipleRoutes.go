package main

import (
	"fmt"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye, World!")
}

func main() {
	mux := http.NewServeMux()

	// Register routes
	mux.Handle("/hello", http.HandlerFunc(helloHandler))
	mux.Handle("/goodbye", http.HandlerFunc(goodbyeHandler))

	// Wrap the entire ServeMux with middleware
	loggedMux := loggingMiddleware(mux)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", loggedMux)
}
