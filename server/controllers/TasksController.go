package controllers

import "net/http"

// TasksController : Controller to hold all routes for tasks
type TasksController struct {
	server *Server
}

// CreateTask : Route to create one task
var createTask = func(writer http.ResponseWriter, request *http.Request) {

}

// GetTaskByID : Route to get a task by Id
var getTaskByID = func(writer http.ResponseWriter, request *http.Request) {

}

// GetAllTasks : Route to get all tasks
var getAllTasks = func(writer http.ResponseWriter, request *http.Request) {

}

// UpdateTask : Route to update a task
var updateTask = func(writer http.ResponseWriter, request *http.Request) {

}

// DeleteTask : Route to delete a task
var deleteTask = func(writer http.ResponseWriter, request *http.Request) {

}
