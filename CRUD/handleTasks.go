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
//variable for waitGroup which is used wait till all the goroutine completed
var wg sync.WaitGroup
// Create a CustomReader for reading user input
var reader = &input.CustomReader{Reader: bufio.NewReader(os.Stdin)}

func HandleTask(w http.ResponseWriter, r *http.Request) {
	//w is the  Responsewriter which is used to write the json and send the http request to the client
	//Responsewriter is like an envelop where u write something and deliver it to the client
	//w is the io.Writer where the json encoded data will be written
	// r contains all the details of the incoming http request like which method get or post etc..
	//http.Request is like a letter with question or request. It basically handle request
	switch r.Method{
	case http.MethodGet:
		//it will return the slice of tasks from the task_manager file 
		tasks:=taskManager.GetTasks()
		//this will set the header  for the response that this response is in json format 
		w.Header().Set("Content-Type","application/json")
		//json NewEncoder is translator which is used to convert from slice to json format in the give program
		//Encode is like writer which writes the tasks in json format
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		//creation of new variable similar to struct Task
		var newTask task.Task
		//This line usually convert the r.body into a slice and then add it to the newtask variable
		//r.body is the request that client send to the program which is in json format
		//Now this json format(New Decoder) is converted to the slice or another data structure then write into the taks(Decode)
		err:=json.NewDecoder(r.Body).Decode(&newTask)
		if err!=nil{
			//it will throw 400 bad request if error else 200 success 
			http.Error(w,"Invalid Input", http.StatusBadRequest)
		}
		error:=taskManager.AddTask(newTask.Title,newTask.DueDate)
		if error!=nil{
			//StatusInternalServerError return the 500 bad request that it is error from server side 
			//err.Error()  return the error but in string fomrat
			//w is the writer where in writes in the body
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		//it writes the status code 
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Task created successfully")

	default:
		//if method an client is rquested is not found
		http.Error(w, "Method Not Allowed",http.StatusMethodNotAllowed)
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
		err := taskManager.CompleteTask(taskID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Task marked as completed")

	case http.MethodDelete:
		// Delete a task
		err := taskManager.DeleteTask(taskID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Task deleted successfully")

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}