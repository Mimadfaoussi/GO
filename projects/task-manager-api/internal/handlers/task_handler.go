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

	if (newTask.Title == "" || newTask.Status == ""){
		http.Error(w, "Title and Status are required fileds.", http.StatusBadRequest)
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

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if (err != nil) {
		http.Error(w, "Wrong Input", http.StatusBadRequest)
		return
	}

	if (updatedTask.Title == "" || updatedTask.Status == ""){
		http.Error(w, "Title and Status are required fileds.", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"UPDATE tasks SET title=$1, description=$2, status=$3 WHERE id=$4",
		updatedTask.Title, updatedTask.Description, updatedTask.Status, updatedTask.ID,
	)

	if (err != nil) {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if (rowsAffected == 0) {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Task updated successfully")
}



func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var TaskToDelete models.Task
	err := json.NewDecoder(r.Body).Decode(&TaskToDelete)
	if (err != nil) {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM tasks WHERE id=$1", TaskToDelete.ID)

	if (err != nil) {
		http.Error(w, "Couldn't delete from tasks", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if (rowsAffected == 0) {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Task deleted successfully")
}


