package service

import (
	"TaskBoard/server/models"
	"TaskBoard/server/repository"
)

// UserService :
type UserService struct {
	config   *models.Config
	userRepo *repository.UserRepository
}

// NewUserService :
func NewUserService(config *models.Config, userRepo *repository.UserRepository) *UserService {
	return &UserService{config: config, userRepo: userRepo}
}
