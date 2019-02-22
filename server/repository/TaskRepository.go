package repository

import (
	"TaskBoard/server/models"
	"TaskBoard/server/util"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// TaskRepository : Struct that holds properties for the TaskRepo
type TaskRepository struct {
	database *gorm.DB
}

// NewTaskRepository : Creates a new instance of the TaskRepository
func NewTaskRepository(database *gorm.DB) *TaskRepository {
	return &TaskRepository{database}
}

// CreateTask : Creates a new task for the user
func (taskRepo *TaskRepository) CreateTask(name string, priority int, uid string) map[string]interface{} {
	tuid, err := GenerateUID()

	if err != nil {
		log.Println("Could not create a UID for this task")
		return util.Message(http.StatusInternalServerError, "Error creating task")
	}

	task := &models.Task{
		Name:     name,
		Priority: priority,
		UID:      uid,
		TID:      tuid,
	}

	taskRepo.database.Create(task)

	gr := util.Message(http.StatusOK, "Successfully added a task")
	return gr
}

// GetTaskByID : Get a task by ID
func (taskRepo *TaskRepository) GetTaskByID(taskid string) map[string]interface{} {
	task := &models.Task{}

	err := taskRepo.database.Table("tasks").Where("t_id = ?", taskid).First(task).Error

	if err != nil {
		log.Println("Error looking up task by Id, ", err.Error())
		return util.Message(http.StatusInternalServerError, "Could not find task")
	}

	gr := util.Message(http.StatusOK, "Success getting task")
	gr["task"] = task
	return gr
}

// GetAllTasks : Get all tasks
func (taskRepo *TaskRepository) GetAllTasks(userid string) map[string]interface{} {
	tasks := make([]*models.Task, 0)

	err := taskRepo.database.Table("tasks").Where("uid = ?", userid).Find(&tasks).Error
	if err != nil {
		log.Println("Error getting all tasks for this user", err.Error())
		return util.Message(http.StatusInternalServerError, "Could not find tasks")
	}

	gr := util.Message(http.StatusOK, "Success found tasks")
	gr["tasks"] = tasks
	return gr
}

// UpdateTaskByID : Update the task
func (taskRepo *TaskRepository) UpdateTaskByID(taskid string, updatedTask models.Task) map[string]interface{} {
	//todo: Will need a clever way of implementing this
	taskRes := taskRepo.GetTaskByID(taskid)
	if taskRes["status"] == 500 {
		log.Println("Could not get task from GetTaskByID")
		return taskRes
	}
	task := taskRes["task"]
	err := taskRepo.database.Model(&task).UpdateColumns(updatedTask).Error

	if err != nil {
		log.Println("Could not update task")
		return util.Message(http.StatusInternalServerError, "Could not update task")
	}

	return util.Message(http.StatusAccepted, "Success updating task")
}

// DeleteTaskByID :
func (taskRepo *TaskRepository) DeleteTaskByID(taskid string) map[string]interface{} {
	return nil
}

// DeleteAllTasks :
func (taskRepo *TaskRepository) DeleteAllTasks() map[string]interface{} {
	return nil
}
