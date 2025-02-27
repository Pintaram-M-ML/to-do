package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)
// Task structure
type Task struct {
	ID    int
	Title string
	Duedate time.Time
}
// CustomReader is a wrapper around bufio.Reader
type CustomReader struct {
	*bufio.Reader
}
//  *bufio.Reader to get user input for task and due date
func (r *CustomReader) inputTheData() (string, time.Time) {
	// Getting task title from user input
	fmt.Println("Enter Your Task: ")
	tsk, _ := r.ReadString('\n')
	tsk = strings.TrimSpace(tsk) // Remove newline

	// Getting due date from user input
	var duedate time.Time
	for {
		fmt.Println("Enter the due date for the task (YYYY-MM-DD): ")
		dueDateStr, _ := r.ReadString('\n')
		dueDateStr = strings.TrimSpace(dueDateStr)

		// Parse the due date
		parsedDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			fmt.Println("Invalid Date format!... Please enter the due date in correct format (YYYY-MM-DD).")
			continue 
		}
		duedate = parsedDate
		break 
	}

	return tsk, duedate
}
func addingTheTask(tasks *[]Task, taskID *int, task string, duedate time.Time) {
	*taskID++
	newTask := Task{
		ID:      *taskID,
		Title:   task,
		Duedate: duedate,
	}
	*tasks = append(*tasks, newTask)
}
// getTheData function to display the list of tasks
func getTheData(tasks []Task) {
	fmt.Println("Your Todo List:")
	for _, task := range tasks {
		// Display each task
		fmt.Printf("ID: %d, Task: %s, Due Date: %s\n", task.ID, task.Title, task.Duedate.Format("2006-01-02"))
	}
}

func main() {
	var tasks []Task
	var taskID int

	// Create a CustomReader to read user input
	reader := &CustomReader{bufio.NewReader(os.Stdin)}
	for {
		// Get task details
		task, duedate := reader.inputTheData()

		// Add task to the list
		addingTheTask(&tasks, &taskID, task, duedate)

		// Ask if the user wants to enter another task
		fmt.Println("Do you want to enter another task? (yes/no)")
		var choice string
		fmt.Scanln(&choice)

		if choice != "yes" {
			break
		}
	}

	// Display all tasks when user is done
	getTheData(tasks)

}

