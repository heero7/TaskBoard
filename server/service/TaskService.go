package service

import (
	"TaskBoard/server/models"
	"TaskBoard/server/repository"
)

// TaskService : Struct that holds properties for the TaskService
type TaskService struct {
	config   *models.Config
	taskRepo *repository.TaskRepository
}

// NewTaskService : Creates a new instance of the TaskService
func NewTaskService(config *models.Config, taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{config: config, taskRepo: taskRepo}
}

// CreateTask : Creates a new task via the repository layer
func (service *TaskService) CreateTask(name string, priority int, uid string) map[string]interface{} {
	return service.taskRepo.CreateTask(name, priority, uid)
}


