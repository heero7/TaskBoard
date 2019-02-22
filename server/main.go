package main

import (
	"TaskBoard/server/controllers"
	"TaskBoard/server/models"
	"TaskBoard/server/repository"
	"TaskBoard/server/service"

	_ "github.com/lib/pq"
)

func main() {
	// without dependency injection

	config := models.NewConfig()

	db, err := repository.ConnectPostgresDatabaseViaGorm(config)
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(&models.User{}, &models.Task{})

	userRepo := repository.NewUserRepository(db)

	taskRepo := repository.NewTaskRepository(db)

	userService := service.NewUserService(config, userRepo)

	taskService := service.NewTaskService(config, taskRepo)

	server := controllers.NewServer(config, userService, taskService)

	server.Start()
}
