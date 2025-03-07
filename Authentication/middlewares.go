package Authentication

import (
	"fmt"
	"net/http"
)

// Middleware for logging requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log incoming requests (method and URL)
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

var users = map[string]string{
	"admin": "password123", // Basic auth example (username: admin, password: password123)
}

// Middleware for basic authentication
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get username and password from basic auth header
		username, password, ok := r.BasicAuth()
		if !ok || users[username] != password {
			// Return Unauthorized status if authentication fails
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
