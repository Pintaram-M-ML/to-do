package input

import (
	"bufio"
	"fmt"
	"strings"
	"time"
	
)

// CustomReader struct for reading user input
type CustomReader struct {
	*bufio.Reader
}

// InputData method to capture user input for task title and due date
func (r *CustomReader) InputData() (string, time.Time, error) {
	// Getting task title from user input
	fmt.Print("Enter your task: ")
	taskTitle, _ := r.ReadString('\n')
	taskTitle = strings.TrimSpace(taskTitle)

	// Getting due date from user input
	var dueDate time.Time
	for {
		fmt.Print("Enter the due date (YYYY-MM-DD): ")
		dueDateStr, _ := r.ReadString('\n')
		dueDateStr = strings.TrimSpace(dueDateStr)

		// Parse the due date
		parsedDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			return "", time.Time{}, fmt.Errorf("invalid date format. Use YYYY-MM-DD")
		}
		if parsedDate.Before(time.Now()) {
			return "", time.Time{}, fmt.Errorf("due date must be in the future")
		}
		dueDate = parsedDate
		break
	}

	return taskTitle, dueDate, nil
}
