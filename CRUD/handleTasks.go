package crud

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"todo-app/internal/input"
	"todo-app/internal/task"
)

// Initialize TaskManager
var taskManager = &task.TaskManager{}

// variable for waitGroup which is used wait till all the goroutine completed
var wg sync.WaitGroup

// Create a CustomReader for reading user input
var reader = &input.CustomReader{Reader: bufio.NewReader(os.Stdin)}

func HandleTask(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        tasks := taskManager.GetTasks()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(tasks)
    case http.MethodPost:
        var newTask task.Task
        err := json.NewDecoder(r.Body).Decode(&newTask)
        if err != nil {
            http.Error(w, "Invalid Input", http.StatusBadRequest)
            return
        }
        fmt.Printf("Decoded task: %+v\n", newTask)

        // Add the task synchronously
        err = taskManager.AddTask(newTask.Title)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        fmt.Fprintf(w, "Task created successfully")
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}


// Handle requests for getting, updating, and deleting a specific task by ID
func HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	// Get the task ID from URL path
	//r.URL wiil return the full url which the client used //eg:=http://www.google.com/xyz/123
	//r.URL.path will return only /xyz/123
	//len will return length and then slice it
	//taskIDstr will return string of id
	taskIDStr := r.URL.Path[len("/tasks/"):]
	//it will convert the string to interger
	//Atoi  =  ASCII to Integer
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Retrieve a specific task by ID
		task, err := taskManager.GetTaskByID(taskID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

	case http.MethodPut:
		// Mark task as completed
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := taskManager.CompleteTask(taskID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
		}()

		fmt.Fprintf(w, "Task marked as completed")

	case http.MethodDelete:
		// Delete a task
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := taskManager.DeleteTask(taskID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
		}()

		fmt.Fprintf(w, "Task deleted successfully")

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	wg.Wait()
}
