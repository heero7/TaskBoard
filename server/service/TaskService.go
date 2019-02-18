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
