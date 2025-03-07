package input

import (
	"bufio"
	"fmt"
	"strings"
)

// CustomReader struct for reading user input
type CustomReader struct {
	*bufio.Reader
}

// InputData method to capture user input for task title and due date
func (r *CustomReader) InputData() (string, error) {
	// Getting task title from user input
	fmt.Print("Enter your task: ")
	taskTitle, _ := r.ReadString('\n')
	taskTitle = strings.TrimSpace(taskTitle)

	// // Getting due date from user input
	// var dueDate time.Time
	// for {
	// 	fmt.Print("Enter the due date (YYYY-MM-DD): ")
	// 	dueDateStr, _ := r.ReadString('\n')
	// 	dueDateStr = strings.TrimSpace(dueDateStr)

	// 	// Parse the due date
	// 	parsedDate, err := time.Parse("2006-01-02", dueDateStr)
	// 	if err != nil {
	// 		fmt.Println("Invalid date format. Please use YYYY-MM-DD format.")
	// 		continue // Ask again if date format is invalid
	// 	}
	// 	if parsedDate.Before(time.Now()) {
	// 		fmt.Println("Due date must be in the future. Please enter a valid date.")
	// 		continue // Ask again if the due date is in the past
	// 	}
	// 	dueDate = parsedDate
	// 	break // Exit the loop when a valid date is entered
	// }

	return taskTitle, nil
}
