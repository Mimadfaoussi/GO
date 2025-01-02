package handlers

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"net/http"
	"task-manager-api/internal/models"
)

// In-memory storage (temporary until database integration)


/*var tasks = []models.Task {
	{ID:1, Title: "Learn Go", Description: "Practice building REST APIS", Status: "In Progress"},
}*/


// Global database instance
var db *sql.DB

func InitDB(database *sql.DB) {
	db = database
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, description, status FROM tasks")
	if (err != nil) {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil{
			http.Error(w, "Error scanning task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}


func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if (err != nil) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = db.Exec("INSERT INTO tasks(title, description, status) VALUES ($1, $2, $3)", newTask.Title, newTask.Description, newTask.Status)
	if (err != nil) {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Task created successfully")
}

// func UpdateTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var updatedTask models.Task
// 	err := json.NewDecoder(r.Body).Decode(&updatedTask)
// 	if (err != nil) {
// 		http.Error(w, "invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	for i, task := range tasks {
// 		if task.ID == updatedTask.ID {
// 			tasks[i] = updatedTask
// 			json.NewEncoder(w).Encode(updatedTask)
// 			return
// 		}
// 	}

// 	http.Error(w, "Task not found", http.StatusNotFound)
// }



// func DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var TaskToDelete models.Task
// 	err := json.NewDecoder(r.Body).Decode(&TaskToDelete)
// 	if (err != nil) {
// 		http.Error(w, "Invalid Input", http.StatusBadRequest)
// 		return
// 	}

// 	for i, task := range tasks {
// 		if task.ID == TaskToDelete.ID {
// 			tasks = append(tasks[:i],tasks[i + 1:]...)
// 			fmt.Fprintln(w, "Task deleted")
// 			return
// 		}
// 	}

// 	http.Error(w, "Task not Found", http.StatusNotFound)
// }


