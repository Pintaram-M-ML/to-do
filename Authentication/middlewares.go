package Authentication

import (
	"fmt"
	"hash"
	"net/http"
)

//loggingmiddleware is like checking the each request  that is like every thing is working fine , which request came , logout etc
// Middleware for logging requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log incoming requests (method and URL)
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		//it will pass the request to the next http handler
		next.ServeHTTP(w, r)
	})
}

var users = map[string]string{
	"admin": "password123", // Basic auth example (username: admin, password: password123)
}
//BasicAuthMiddleware is like checking each request made by customer is authorised or not
// Middleware for basic authentication
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get username and password from basic auth header
		//It’s used to extract the basic authentication credentials (username and password) from the HTTP request headers.
		username, password, ok := r.BasicAuth()
		if !ok || users[username] != password {
			// Return Unauthorized status if authentication fails
			//fmt.Println("admin" ,username , "\n password",password )
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		//it will pass the http request to the next handler
		next.ServeHTTP(w, r)
	})
}
