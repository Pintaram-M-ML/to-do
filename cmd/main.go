package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"todo-app/internal/input"
	"todo-app/internal/task"
     
)

func main() {
	// Initialize TaskManager
	taskManager := &task.TaskManager{}

	// Create a CustomReader for reading user input
	reader := &input.CustomReader{Reader: bufio.NewReader(os.Stdin)}


	for {
		// Show options
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add a task")
		fmt.Println("2. Delete a task")
		fmt.Println("3. Mark a task as completed")
		fmt.Println("4. Display tasks")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Add task
			taskTitle, dueDate, err := reader.InputData()
			if err != nil {
				log.Println("Error:", err)
			} else {	
				err := taskManager.AddTask(taskTitle, dueDate)
				if err != nil {
					log.Println("Error adding task:", err)
				} else {
					fmt.Println("Task added successfully!")
				}			
				
			}

		case 2:
			// Delete task
			fmt.Print("Enter the Task ID to delete: ")
			var taskIDToDelete int
			fmt.Scanln(&taskIDToDelete)
				err := taskManager.DeleteTask(taskIDToDelete)
				if err != nil {
					log.Println("Error:", err)
				} else {
	
					fmt.Println("Task deleted successfully!")
				}		

		case 3:
			// Mark task as completed
			fmt.Print("Enter the Task ID to mark as completed: ")
			var taskIDToComplete int
			fmt.Scanln(&taskIDToComplete)
				err := taskManager.CompleteTask(taskIDToComplete)
			if err != nil {
				log.Println("Error:", err)
			} else {
				fmt.Println("Task marked as completed!")
			}		

		case 4:
			// Display tasks
			tasks := taskManager.GetTasks()
			task.PrintTasks(tasks)

		case 5:
			// Exit
			fmt.Println("Thank you for using the To-Do list application!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
