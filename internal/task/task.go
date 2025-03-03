package task

import "time"

// Task structure to represent each task
type Task struct {
	ID        int
	Title     string
	DueDate   time.Time
	Completed bool
}
