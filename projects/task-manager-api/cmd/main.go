package main

import (
	"fmt"
	"net/http"
	"task-manager-api/internal/handlers"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to task manager API")
}

func main() {


	http.HandleFunc("/", homeHandler)


	http.HandleFunc("/tasks", handlers.GetTasks)
	http.HandleFunc("/tasks/create", handlers.CreateTask)
	http.HandleFunc("/tasks/update", handlers.UpdateTask)
	http.HandleFunc("/tasks/delete", handlers.DeleteTask)


	port := ":8000"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting the server")
	}
}

