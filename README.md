
# To-Do App with Authentication

This is a simple To-Do application written in Go, which implements task management with features like adding, updating, deleting, and marking tasks as completed. The app uses basic authentication and logging middleware to secure the API.

## Features

- User Authentication: Basic Authentication using a username and password (currently hardcoded in the app).
Task Management: 
  - Add tasks
  - View tasks
  - Complete tasks
  - Delete tasks
- Logging Middleware: Logs incoming requests for debugging and monitoring.
- Concurrency: Uses Goroutines and WaitGroups to handle tasks asynchronously.

## Getting Started with To Do application

### Prerequisites

Before running the application, ensure you have Go installed on your system.

1. Download and Install Go (https://golang.org/doc/install) (if you don't have it already).

2. Clone the repository to your local machine:

   git clone  https://github.com/Pintaram-M-ML/to-do.git
   cd todo-app
   cd cmd

### Running the Application

1. In the project folder, run the following command to start the server:

   go run main.go

   The server will start on `http://localhost:8080`.

2. Open your browser or use an API testing tool like Postman to interact with the API.

### API Endpoints

- POST `/tasks`: Create a new task.

  - Request body: 
    
    {
      "title": "Task Title"
    }
    
  - Response: A JSON object containing the task ID and title.

- GET `/tasks`:  Retrieve a list of all tasks.
  - Response: A JSON array of tasks.

- GET  `/tasks/{taskID}`: Retrieve a specific task by its ID.
  - Response: A JSON object containing the task details.

- PUT  `/tasks/{taskID}`: Mark a task as completed.
  - Response: Status message indicating success.

- DELETE `/tasks/{taskID}`: Delete a task by its ID.
  - Response: Status message indicating success.


## Concurrency

The app uses goroutines for handling the creation, deletion, and completion of tasks asynchronously. The sync.WaitGroup is used to ensure that all goroutines finish before sending the response to the client.

### Example of Goroutines in Use

- When a task is added, the `POST /tasks` endpoint spawns a goroutine to handle the task creation asynchronously.
- The `WaitGroup` is used to ensure that the server waits for the task creation process to complete before responding to the client.

## Code Structure

- main.go: The entry point for the application. Sets up routes and middleware.
- Authentication: Contains the middleware for logging and basic authentication.
- CRUD: Handles the business logic for task management, such as adding, deleting, and updating tasks.
- Task: Contains the task structure and methods for managing tasks.

<!-- ## Future Improvements

- Implement a database to store tasks instead of in-memory storage.
- Add user registration and store credentials securely (e.g., using bcrypt for password hashing).
- Expand the API to include more task details (e.g., due dates, descriptions).
- Add more detailed logging and error handling. -->
