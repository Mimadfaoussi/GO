package main

import (
	"fmt"
	"net/http"
	"task-manager-api/internal/handlers"
	"task-manager-api/internal/db"
	"task-manager-api/internal/middleware"
	"github.com/rs/cors"

)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to task manager API")
}

func main() {

	database := db.ConnectDB()
	defer database.Close()

	handlers.InitDB(database)
	// http.HandleFunc("/", homeHandler)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", handlers.GetTasks)
	mux.HandleFunc("/tasks/create", handlers.CreateTask)
	mux.HandleFunc("/tasks/update", handlers.UpdateTask)
	mux.HandleFunc("/tasks/delete", handlers.DeleteTask)

	loggedMux := middleware.LoggingMiddleware(mux)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(loggedMux)


	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, corsMiddleware)
	if err != nil {
	 	fmt.Println("Error starting the server")
	}
}

