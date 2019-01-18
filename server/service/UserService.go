package service

import (
	"TaskBoard/server/models"
	"TaskBoard/server/repository"
)

// UserService : Struct that holds the properties for the UserService
type UserService struct {
	config   *models.Config
	userRepo *repository.UserRepository
}

// NewUserService : Creates a new instance of the UserService
func NewUserService(config *models.Config, userRepo *repository.UserRepository) *UserService {
	return &UserService{config: config, userRepo: userRepo}
}

// CreateUser : Service to create a user via the UserRepository
func (service *UserService) CreateUser(email string, password string) map[string]interface{} {
	return service.userRepo.CreateUser(email, password)
}

// Authenticate : Will attempt to login a character
func (service *UserService) Authenticate(email string, password string) error {
	return nil
}
