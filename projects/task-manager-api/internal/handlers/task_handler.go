package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"task-manager-api/internal/models"
)

// In-memory storage (temporary until database integration)


var tasks = []models.Task { 
	{ID:1, Title: "Learn Go", Description: "Practice building REST APIS", Status: "In Progress"},
}


func GetTasks(w http.RequestWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}


func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTask models.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if (err != nil) {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}












