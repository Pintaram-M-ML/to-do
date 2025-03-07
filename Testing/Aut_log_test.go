package Testing

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-app/Authentication"
)

func testBasicAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func TestBasicAuthMiddleware(t *testing.T) {
	handler := Authentication.BasicAuthMiddleware(http.HandlerFunc(testBasicAuthHandler))

	// Create a valid request with the basic authentication
	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("admin", "password123") // valid credentials

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check that the status is OK (200)
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %v", rr.Code)
	}

	// Create an invalid request with incorrect credentials
	req.SetBasicAuth("admin", "wrongpassword") // invalid credentials

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check that the status is Unauthorized (401)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %v", rr.Code)
	}
}
