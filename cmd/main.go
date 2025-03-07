package main

import (
	"fmt"
	"net/http"
	"todo-app/Authentication"
	crud "todo-app/CRUD"
)

func main() {
	// http.HandleFunc("/tasks", crud.HandleTask)
	// http.HandleFunc("/tasks/", crud.HandleTaskByID)
	// Wrap the handlers with logging middleware
	// Wrap handlers with logging and authentication middleware
	http.Handle("/tasks", Authentication.BasicAuthMiddleware(Authentication.LoggingMiddleware(http.HandlerFunc(crud.HandleTask))))
	http.Handle("/tasks/", Authentication.BasicAuthMiddleware(Authentication.LoggingMiddleware(http.HandlerFunc(crud.HandleTaskByID))))
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	//instead of running in server
	//crud.User_Input()
}
