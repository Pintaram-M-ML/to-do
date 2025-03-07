package Testing

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-app/CRUD"
	"todo-app/internal/task"
)
func TestHandleTask(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(crud.HandleTask))
    defer ts.Close()

    // Add a task with POST
    postResp, err := http.Post(ts.URL+"/tasks", "application/json", strings.NewReader(`{"Title":"Test Task"}`))
    if err != nil {
        t.Fatal(err)
    }
    defer postResp.Body.Close()

    // Ensure task was created successfully
    if postResp.StatusCode != http.StatusCreated {
        t.Fatalf("expected status 201 Created, got %v", postResp.Status)
    }

    // Send GET request to retrieve tasks
    getResp, err := http.Get(ts.URL + "/tasks")
    if err != nil {
        t.Fatal(err)
    }
    defer getResp.Body.Close()

    // Check if the response status is OK
    if getResp.StatusCode != http.StatusOK {
        t.Fatalf("expected status OK, got %v", getResp.Status)
    }

    // Decode the response body
    var tasks []task.Task
    if err := json.NewDecoder(getResp.Body).Decode(&tasks); err != nil {
        t.Fatal(err)
    }

    // Check if the returned task matches the one created
    if len(tasks) != 1 {
        t.Fatalf("expected 1 task, got %d", len(tasks))
    }

    if tasks[0].Title != "Test Task" {
        t.Errorf("expected task title 'Test Task', got %v", tasks[0].Title)
    }
}
