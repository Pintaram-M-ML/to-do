package Testing

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-app/Authentication"
)

// A simple handler to test logging middleware
func testLoggingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func TestLoggingMiddleware(t *testing.T) {
	// Wrap the handler with the LoggingMiddleware
	handler := Authentication.LoggingMiddleware(http.HandlerFunc(testLoggingHandler))

	// Create a new HTTP request to test the middleware
	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler with the request and recorder
	handler.ServeHTTP(rr, req)

	// Assert the status code is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %v", rr.Code)
	}
}
