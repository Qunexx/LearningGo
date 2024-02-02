package main

import (
	"encoding/json"
	"fmt"
	"myapp/TaskSchedulerMicroservice"
	"net/http"
)

func main() {

	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/task/", taskHandler)
	fmt.Printf("Сервер запущен")
	http.ListenAndServe(":8080", nil)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		list := TaskSchedulerMicroservice.GetTasks()
		json.NewEncoder(w).Encode(list)
	} else if r.Method == "POST" {
		var task TaskSchedulerMicroservice.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdTask := TaskSchedulerMicroservice.CreateTask(task.ID, task.Description)
		json.NewEncoder(w).Encode(createdTask)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/task/"):]
	if r.Method == "PUT" {
		var task TaskSchedulerMicroservice.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedTask, err := TaskSchedulerMicroservice.UpdateTask(id, task.Completed)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(updatedTask)
	} else if r.Method == "DELETE" {
		if err := TaskSchedulerMicroservice.DeleteTask(id); err != nil {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
