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

	db, err := repository.ConnectPostgresDatabase(config)

	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(config, userRepo)

	server := NewServer(config, userService)

	server.Start()
}
