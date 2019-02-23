package controllers

import (
	"TaskBoard/server/models"
	"TaskBoard/server/service"
	"TaskBoard/server/util"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	taskID := mux.Vars(request)["id"]
	res := tc.taskService.GetTaskByID(taskID)
	util.Respond(writer, res)
}

// GetAllTasks : Route to get all tasks
func (tc *TasksController) getAllTasks(writer http.ResponseWriter, request *http.Request) {
	userID := request.Context().Value("user").(string)
	res := tc.taskService.GetAllTasks(userID)
	util.Respond(writer, res)
}

// UpdateTask : Route to update a task
func (tc *TasksController) updateTask(writer http.ResponseWriter, request *http.Request) {

}

// DeleteTask : Route to delete a task
func (tc *TasksController) deleteTask(writer http.ResponseWriter, request *http.Request) {
	taskID := mux.Vars(request)["id"]
	res := tc.taskService.DeleteTask(taskID)
	util.Respond(writer, res)
}
