package repository

import (
	"github.com/jinzhu/gorm"
)

// TaskRepository : Struct that holds properties for the TaskRepo
type TaskRepository struct {
	database *gorm.DB
}

// NewTaskRepository : Creates a new instance of the TaskRepository
func NewTaskRepository(database *gorm.DB) *TaskRepository {
	return nil
}
