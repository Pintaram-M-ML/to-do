package Testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	crud "todo-app/CRUD"
	"todo-app/internal/task"
)

// TestAddTask checks if a new task is added successfully.
//  //TestAddTask checks if a new task is added successfully.
func TestAddTask(t *testing.T) {
	// Create a new task as JSON
	taskJSON := `{"Title": "Test Task"}`
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(taskJSON)))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("admin", "password123") // Add basic auth if necessary

	// Capture the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crud.HandleTask)

	// Perform the POST request
	handler.ServeHTTP(rr, req)

	// Assert that the status code is 201 Created
	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201 Created, got %v", rr.Code)
	}

	// Decode the response body to check if the task was added
	var response map[string]string
	fmt.Println(rr.Body)
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	// Ensure the response has the correct message
	expected := "Task created successfully"
	if response["message"] != expected {
		t.Errorf("expected message %v, got %v", expected, response["message"])
	}
}

//get task
func TestGetTasks(t *testing.T) {
	// Create a new HTTP request to fetch tasks
	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("admin", "password123") // Add auth if necessary

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crud.HandleTask)

	// Perform the request
	handler.ServeHTTP(rr, req)

	// Check the response status
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", rr.Code)
	}

	// Decode the response body
	var tasks []task.Task
	if err := json.NewDecoder(rr.Body).Decode(&tasks); err != nil {
		t.Fatal(err)
	}

	// Assert that tasks were returned
	if len(tasks) == 0 {
		t.Errorf("expected tasks, but got none")
	}
}

//DELETE
func TestDeleteTask(t *testing.T) {
	// Create a new task first
	taskJSON := `{"Title": "Test Task"}`
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(taskJSON)))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("admin", "password123")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crud.HandleTask)

	// Perform the request to create the task
	handler.ServeHTTP(rr, req)

	// Now, create a request to delete the task
	deleteReq, err := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	deleteReq.SetBasicAuth("admin", "password123")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(crud.HandleTaskByID)

	// Perform the delete request
	handler.ServeHTTP(rr, deleteReq)

	// Check that the task was deleted successfully
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", rr.Code)
	}
}
 
//PUT TASK
func TestUpdateTask(t *testing.T) {
	// Create a new task first
	taskJSON := `{"Title": "Test Task"}`
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(taskJSON)))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("admin", "password123")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(crud.HandleTask)

	// Perform the request to create the task
	handler.ServeHTTP(rr, req)

	// Now, create a request to mark the task as completed
	updateReq, err := http.NewRequest(http.MethodPut, "/tasks/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	updateReq.SetBasicAuth("admin", "password123")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(crud.HandleTaskByID)

	// Perform the update request (mark as completed)
	handler.ServeHTTP(rr, updateReq)

	// Check that the task was updated
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", rr.Code)
	}
}
