package middleware

import (
	"fmt"
	"net/http"
	"time"
)


// the logging middleware logs each request

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[%s] %s %s\n", r.Method, r.URL.Path, start.Format(time.RFC3339))
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("Completed in : %v\n", duration)
	})
}


