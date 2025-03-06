package task

import (
	"fmt"
	"sync"
)

// TaskManager struct to hold tasks
type TaskManager struct {
	tasks   []Task
	taskID  int
	mut sync.Mutex
}

// AddTask adds a task to the TaskManager
func (tm *TaskManager) AddTask(title string) error {
	tm.mut.Lock()
	defer tm.mut.Unlock()
	tm.taskID++
	newTask := Task{
		ID:      tm.taskID,
		Title:   title,
	
	}
	tm.tasks = append(tm.tasks, newTask)
	return nil
}

// GetTasks returns all tasks
func (tm *TaskManager) GetTasks() []Task {
	tm.mut.Lock()
	defer tm.mut.Unlock()
	return tm.tasks
}

// CompleteTask marks a task as completed
func (tm *TaskManager) CompleteTask(taskID int) error {
	tm.mut.Lock()
	defer tm.mut.Unlock()
	for i, task := range tm.tasks {
		if task.ID == taskID {
			tm.tasks[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", taskID)
}
// GetTaskByID returns a task by its ID
func (tm *TaskManager) GetTaskByID(id int) (*Task, error) {
	tm.mut.Lock()
	defer tm.mut.Unlock()

	for _, task := range tm.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}


// DeleteTask deletes a task from the task list
func (tm *TaskManager) DeleteTask(taskID int) error {
	tm.mut.Lock()
	defer tm.mut.Unlock()
	for i, task := range tm.tasks {
		if task.ID == taskID {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", taskID)
}

// PrintTasks prints all tasks in a formatted table
func PrintTasks(tasks []Task) {
	fmt.Println("\nTask List:")
	fmt.Printf("| %-4s | %-25s | %-15s |\n", "ID", "Task", "Status")
	fmt.Println("|------|---------------------------|--------------|-----------------|")
	for _, task := range tasks {
		status := "Not Completed"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("| %-4d | %-25s | %-12s | %-15s |\n", task.ID, task.Title,status)
	}
}
