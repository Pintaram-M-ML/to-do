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
	completed bool
}
// CustomReader is a wrapper around bufio.Reader
type CustomReader struct {
	*bufio.Reader
}
//  *bufio.Reader to get user input for task and due date
func (r *CustomReader) inputTheData() (string, time.Time) {
	// Getting task title from user input
	fmt.Print("Enter Your Task: ")
	tsk, _ := r.ReadString('\n')
	tsk = strings.TrimSpace(tsk) // Remove newline

	// Getting due date from user input
	var duedate time.Time
	for {
		fmt.Print("Enter the due date for the task (YYYY-MM-DD): ")
		dueDateStr, _ := r.ReadString('\n')
		dueDateStr = strings.TrimSpace(dueDateStr)

		// Parse the due date
		parsedDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			fmt.Println("Invalid Date format!... Please enter the due date in correct format (YYYY-MM-DD).")
			continue 
		}
		if parsedDate.Before(time.Now()) {
			fmt.Println("The due date must be greater then current date. Please enter a valid due date (in the future).")
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
// getTheData function to display the list of tasks in a table format
func getTheData(tasks []Task) {
	println()
	// Print header
	fmt.Printf("| %-4s | %-25s | %-12s | %-15s |\n", "ID", "Task", "Due Date", "Status")
	fmt.Println("|------|---------------------------|--------------|-----------------|")

	// Print each task
	for _, task := range tasks {
		completedStatus := "Not Completed"
		if task.completed {
			completedStatus = "Completed"
		}
		// Print each task's details with formatting
		fmt.Printf("| %-4d | %-25s | %-12s | %-15s |\n", task.ID, task.Title, task.Duedate.Format("2006-01-02"), completedStatus)
	}
}

//mark the task as completed when user enter as completed
func completedTheTask(task *[]Task, taskId int) error {
	for i, tasks := range *task{
		if taskId == tasks.ID{
			(*task)[i].completed = true
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found please check with id ", taskId)
} 
 
//function to delete the task on user command
func deleteTheTask(task *[]Task,taskId int) error  {
	for i, tasks := range *task{
		if taskId == tasks.ID{
			*task = append((*task)[:i], (*task)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found please check with id ", taskId)
}
func main() {
	var tasks []Task
	var taskID int

	// Create a CustomReader to read user input
	reader := &CustomReader{bufio.NewReader(os.Stdin)}
	
    for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add a task to the application")
		fmt.Println("2. Delete a task from the application")
		fmt.Println("3. Mark a task as completed")
		fmt.Println("4. Display the Task the ids and their status")
		fmt.Println("5. Exit from the application\n")
	
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Get task details
			task, duedate := reader.inputTheData()

			// Add task to the list
			addingTheTask(&tasks, &taskID, task, duedate)

		case 2:
			// Ask for the task ID to delete
			fmt.Println("Enter the Task ID to delete:")
			var taskIDToDelete int
			fmt.Scanln(&taskIDToDelete)

			// Try deleting the task
			err := deleteTheTask(&tasks, taskIDToDelete)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted successfully!")
			}
		case 3:
			// Ask for the task ID to mark as completed
			fmt.Println("Enter the Task ID to mark as completed:")
			var taskIDToComplete int
			fmt.Scanln(&taskIDToComplete)

			// Try marking the task as completed
			err := completedTheTask(&tasks, taskIDToComplete)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task marked as completed!")
				
			}
		case 4:
			//to display the  task with updated status
			fmt.Println("Updated  To-Do List with task id and status")
			getTheData(tasks)
		case 5:
			fmt.Println("Thank You for choosing us")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}

}

