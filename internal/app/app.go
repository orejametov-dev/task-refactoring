package app

import (
	"task/internal/domain"
	taskService "task/internal/services/tasks"
	"time"
)

const (
	MAX_EXECUTION_COUNT           = 10
	MAX_EXECUTION_TIME            = time.Duration(20) * time.Second
	ARTIFICAIAL_APPLICATION_PAUSE = time.Duration(2) * time.Second
)

// Run initializes whole application.
func Run() {
	tasks := make(chan domain.Task, MAX_EXECUTION_COUNT)
	completedTasks := make(chan domain.Task, MAX_EXECUTION_COUNT)

	go taskService.Creator(tasks)
	go taskService.Worker(tasks, completedTasks)
	taskService.Print(completedTasks)
}
