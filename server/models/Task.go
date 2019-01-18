package models

import (
	"github.com/jinzhu/gorm"
)

// Task : Model object that maps to the task object.
type Task struct {
	gorm.Model
	Name     string
	Priority int
	TID      string
	UID      string
}
