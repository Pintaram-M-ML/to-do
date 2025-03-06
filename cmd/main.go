package main

import (
	"fmt"
	"net/http"
	"todo-app/CRUD"
)

func main() {
	http.HandleFunc("/tasks", crud.HandleTask)
	http.HandleFunc("/tasks/", crud.HandleTaskByID)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	//instead of running in server
	//crud.User_Input()
}
