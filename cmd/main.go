package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"todo-app/internal/input"
	"todo-app/internal/task"

)

func main() {
	// Initialize TaskManager
	taskManager := &task.TaskManager{}
	//variable for waitGroup which is used wait till all the goroutine completed
    var wg sync.WaitGroup
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
				wg.Add(1)
				go func() {
					defer wg.Done()
					err := taskManager.AddTask(taskTitle, dueDate)
				if err != nil {
					log.Println("Error adding task:", err)
				} 
				}()
				
			}

		case 2:
			// Delete task
			fmt.Print("Enter the Task ID to delete: ")
			var taskIDToDelete int
			fmt.Scanln(&taskIDToDelete)
			//channel can be used instead of wait.group but main thread will wait untill the goroutine 
			//thread task is completed so we can used wait.group
			//In this code we don't need any channel creation 
			//basically channel is used for communication between the goroutine
			//channel creation
			//ch :=make(chan string)
			wg.Add(1)
			go func() {
				defer wg.Done()
				//passing the channel message as "done" from goroutine  thread to main thread  
				//ch<-"done"
				err := taskManager.DeleteTask(taskIDToDelete)
				if err != nil {
					log.Println("Error:", err)
				}
			}()
			// receiveing the "done" message from goroutine thread to main thread
			//<-ch
					
		case 3:
			// Mark task as completed
			fmt.Print("Enter the Task ID to mark as completed: ")
			var taskIDToComplete int
			fmt.Scanln(&taskIDToComplete)
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := taskManager.CompleteTask(taskIDToComplete)
			if err != nil {
				log.Println("Error:", err)
			} 
			}()
			

		case 4:
			// Display tasks
			tasks := taskManager.GetTasks()
			task.PrintTasks(tasks)

		case 5:
			// Exit
			wg.Wait()
			fmt.Println("Thank you for using the To-Do list application!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
