package main

import (
	"fmt"
	"net/http"
	"task-manager-api/internal/handlers"
	"task-manager-api/internal/db"
	"task-manager-api/internal/middleware"
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
	mux.http.HandleFunc("/tasks", handlers.GetTasks)
	mux.http.HandleFunc("/tasks/create", handlers.CreateTask)
	mux.http.HandleFunc("/tasks/update", handlers.UpdateTask)
	mux.http.HandleFunc("/tasks/delete", handlers.DeleteTask)

	loggedMux := middleware.LoggingMiddleware(mux)

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
	 	fmt.Println("Error starting the server")
	}
}

