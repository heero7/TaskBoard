package controllers

import (
	"TaskBoard/server/models"
	"TaskBoard/server/service"
	"TaskBoard/server/util"
	"encoding/json"
	"fmt"
	"net/http"
)

// TasksController : Controller to hold all routes for tasks
type TasksController struct {
	taskService *service.TaskService
}

func initTasksController(ts *service.TaskService) *TasksController {
	return &TasksController{
		taskService: ts,
	}
}

// CreateTask : Route to create one task
func (tc *TasksController) createTask(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(string)
	fmt.Println(user)
	task := &models.Task{}

	err := json.NewDecoder(request.Body).Decode(task)

	if err != nil {
		util.Respond(writer, util.Message(http.StatusInternalServerError, "Error decoding request body"))
		return
	}

	task.UID = user //strconv.Itoa(user)

	res := tc.taskService.CreateTask(task.Name, task.Priority, task.UID)
	util.Respond(writer, res)
}

// GetTaskByID : Route to get a task by Id
func (tc *TasksController) getTaskByID(writer http.ResponseWriter, request *http.Request) {

}

// GetAllTasks : Route to get all tasks
func (tc *TasksController) getAllTasks(writer http.ResponseWriter, request *http.Request) {

}

// UpdateTask : Route to update a task
func (tc *TasksController) updateTask(writer http.ResponseWriter, request *http.Request) {

}

// DeleteTask : Route to delete a task
func (tc *TasksController) deleteTask(writer http.ResponseWriter, request *http.Request) {

}
