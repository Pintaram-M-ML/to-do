package task

import "time"

// TaskManagerInterface defines the required methods for task management
type TaskManagerInterface interface {
	AddTask(title string, dueDate time.Time) error
	DeleteTask(taskID int) error
	CompleteTask(taskID int) error
	GetTasks() []Task
}
