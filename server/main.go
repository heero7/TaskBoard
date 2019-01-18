package main

import (
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

	db.Debug().AutoMigrate(&models.User{}, &models.User{})

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(config, userRepo)

	server := NewServer(config, userService)

	server.Start()
}
