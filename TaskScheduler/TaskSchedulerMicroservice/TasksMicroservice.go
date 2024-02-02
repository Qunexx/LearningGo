package TaskSchedulerMicroservice

import (
	"errors"
	"sync"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var (
	tasks = make(map[string]Task)
	mu    sync.Mutex
)

func CreateTask(id, description string) Task {
	mu.Lock()
	defer mu.Unlock()
	task := Task{ID: id, Description: description, Completed: false}
	tasks[id] = task
	return task
}

func GetTasks() []Task {
	mu.Lock()
	defer mu.Unlock()
	list := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		list = append(list, task)
	}
	return list
}

func UpdateTask(id string, completed bool) (Task, error) {
	mu.Lock()
	defer mu.Unlock()
	if task, ok := tasks[id]; ok {
		task.Completed = completed
		tasks[id] = task
		return task, nil
	}
	return Task{}, errors.New("task not found")
}

func DeleteTask(id string) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := tasks[id]; ok {
		delete(tasks, id)
		return nil
	}
	return errors.New("task not found")
}
