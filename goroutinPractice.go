// // // package main

// // // import (
// // // 	"fmt"
// // // 	"sync"
// // // )

// // // var counter = 0
// // // var mu sync.Mutex  // Declare a mutex

// // // func increment() {
// // // 	mu.Lock()         // Lock the mutex to ensure only one goroutine can access the counter at a time
// // // 	counter++
// // // 	mu.Unlock()       // Unlock the mutex to allow other goroutines to access the counter
// // // }

// // // func main() {
// // // 	// Launch 1000 goroutines to increment the counter
// // // 	var wg sync.WaitGroup // A wait group to wait for all goroutines to finish

// // // 	for i := 0; i < 1000; i++ {
// // // 		wg.Add(1)  // Add 1 to the wait group counter
// // // 		go func() {
// // // 			defer wg.Done() // Decrease the counter when the goroutine is done
// // // 			increment()
// // // 		}()
// // // 	}

// // // 	// Wait for all goroutines to finish
// // // 	wg.Wait()

// // // 	fmt.Println("Counter value:", counter)
// // // }



// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"sync"
// 	"time"
// )

// // Task structure
// type Task struct {
// 	ID    int
// 	Title string
// 	Duedate time.Time
// 	completed bool
// }
// type toDoMangager interface{
//   inputTheData()(string ,time.Time)
//   deleteTheTask(taskID int)error
//   addingTheTask(task string , dueDate time.Time)error
//   completedTheTask(taskID int)
//   getTheData()
  
// }
// //this structure implements the interface 
// type TaskManager struct{
//  tasks []Task
//  taskID int
//  mu sync.Mutex
// }
// //result for goroutin and channel
// type TaskResult struct {
// 	Action string
// 	Success bool
// 	Message string
// }
// // CustomReader is a wrapper around bufio.Reader
// type CustomReader struct {
// 	*bufio.Reader
// }
// //  *bufio.Reader to get user input for task and due date
// func (r *CustomReader) inputTheData() (string, time.Time) {
// 	// Getting task title from user input
// 	fmt.Print("Enter Your Task: ")
// 	tsk, _ := r.ReadString('\n')
// 	tsk = strings.TrimSpace(tsk) // Remove newline

// 	// Getting due date from user input
// 	var duedate time.Time
// 	for {
// 		fmt.Print("Enter the due date for the task (YYYY-MM-DD): ")
// 		dueDateStr, _ := r.ReadString('\n')
// 		dueDateStr = strings.TrimSpace(dueDateStr)

// 		// Parse the due date
// 		parsedDate, err := time.Parse("2006-01-02", dueDateStr)
// 		if err != nil {
// 			fmt.Println("Invalid Date format!... Please enter the due date in correct format (YYYY-MM-DD).")
// 			continue 
// 		}
// 		if parsedDate.Before(time.Now()) {
// 			fmt.Println("The due date must be greater then current date. Please enter a valid due date (in the future).")
// 			continue
// 		}

// 		duedate = parsedDate
// 		break 
// 	}

// 	return tsk, duedate
// }
// func (tm *TaskManager) addingTheTask(task string, duedate time.Time) {
// 	tm.mu.Lock()
// 	defer tm.mu.Unlock()
// 	tm.taskID++
// 	newTask := Task{
// 		ID:      tm.taskID,
// 		Title:   task,
// 		Duedate: duedate,
// 	}
// 	tm.tasks = append(tm.tasks, newTask)
// }
// // getTheData function to display the list of tasks in a table format
// func (tm *TaskManager) getTheData() {
// 	tm.mu.Lock()
// 	defer tm.mu.Unlock()
// 	println()
// 	// Print header
// 	fmt.Printf("| %-4s | %-25s | %-12s | %-15s |\n", "ID", "Task", "Due Date", "Status")
// 	fmt.Println("|------|---------------------------|--------------|-----------------|")

// 	// Print each task
// 	for _, task := range tm.tasks {
// 		completedStatus := "Not Completed"
// 		if task.completed {
// 			completedStatus = "Completed"
// 		}
// 		// Print each task's details with formatting
// 		fmt.Printf("| %-4d | %-25s | %-12s | %-15s |\n", task.ID, task.Title, task.Duedate.Format("2006-01-02"), completedStatus)
// 	}
// }

// //mark the task as completed when user enter as completed
// func (tm *TaskManager)completedTheTask(taskId int) error {
// 	tm.mu.Lock()
// 	defer tm.mu.Unlock()
// 	for i, tasks := range tm.tasks{
// 		if taskId == tasks.ID{
// 			(tm.tasks)[i].completed = true
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("task with ID %d not found please check with id ", taskId)
// } 
 
// //function to delete the task on user command
// func (tm *TaskManager) deleteTheTask(taskId int) error  {
// 	tm.mu.Lock()
// 	defer tm.mu.Unlock()
// 	for i, tasks := range tm.tasks{
// 		if taskId == tasks.ID{
// 			tm.tasks = append((tm.tasks)[:i], (tm.tasks)[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("task with ID %d not found please check with id ", taskId)
// }
// func main() {
// 	taskManager := &TaskManager{}
// 	// Create channels to communicate results from goroutines
// 	addTaskCh := make(chan TaskResult)
// 	deleteTaskCh := make(chan TaskResult)
// 	completeTaskCh := make(chan TaskResult)
// 	// Create a CustomReader to read user input
// 	reader := &CustomReader{bufio.NewReader(os.Stdin)}
	
//     for {
// 		fmt.Println("\nChoose an option:")
// 		fmt.Println("1. Add a task to the application")
// 		fmt.Println("2. Delete a task from the application")
// 		fmt.Println("3. Mark a task as completed")
// 		fmt.Println("4. Display the Task the ids and their status")
// 		fmt.Println("5. Exit from the application")
// 		fmt.Println()
	
// 		var choice int
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
// 			go func() {
// 				// Get task details
// 			task, duedate := reader.inputTheData()

// 			// Add task to the list
// 			taskManager.addingTheTask(task, duedate)
// 			//channel
// 			addTaskCh <- TaskResult{"add", true, "Task added successfully!"}
// 			}()
			

// 		case 2:
// 			go func() {
// 				// Ask for the task ID to delete
// 			fmt.Println("Enter the Task ID to delete:")
// 			var taskIDToDelete int
// 			fmt.Scanln(&taskIDToDelete)

// 			// Try deleting the task
// 			err := taskManager.deleteTheTask(taskIDToDelete)
// 			if err != nil {
// 				deleteTaskCh <- TaskResult{"delete", false, err.Error()}
// 			} else {
// 				deleteTaskCh <- TaskResult{"delete", true, "Task deleted successfully!"}
// 			}
// 			}()
			
// 		case 3:
// 			go func() {
// 			// Ask for the task ID to mark as completed
// 			fmt.Println("Enter the Task ID to mark as completed:")
// 			var taskIDToComplete int
// 			fmt.Scanln(&taskIDToComplete)
// 			// Try marking the task as completed
// 			err := taskManager.completedTheTask(taskIDToComplete)
// 			if err != nil {
// 				completeTaskCh <- TaskResult{"complete", false, err.Error()}
// 			} else {
// 				completeTaskCh <- TaskResult{"complete", true, "Task marked as completed!"}
// 			}	
// 			}()

// 		case 4:
// 			//to display the  task with updated status
// 			fmt.Println("Updated  To-Do List with task id and status")
// 			taskManager.getTheData()
// 		case 5:
// 			fmt.Println("Thank You for choosing us")
// 			return
// 		default:
// 			fmt.Println("Invalid choice, please try again.")
// 		}
// 	}
// 	// // Handling results from channels
// 	select {
// 	case result := <-addTaskCh:
// 		fmt.Println(result.Message)
// 	case result := <-deleteTaskCh:
// 		fmt.Println(result.Message)
// 	case result := <-completeTaskCh:
// 		fmt.Println(result.Message)
// 	}

// }

