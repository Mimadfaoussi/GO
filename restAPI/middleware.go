package main

import (
	"fmt"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Method: %s, URL: %s, Time: %s\n", r.Method, r.URL.Path, start.Format(time.RFC3339))
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}


func main() {
	http.Handle("/hello", loggingMiddleware(http.HandlerFunc(helloHandler)))
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
