package main

import (
	"TaskBoard/server/models"
	"TaskBoard/server/service"
	"encoding/json"
	"fmt"
	"net/http"
)

// Server :
type Server struct {
	config      *models.Config
	userService *service.UserService
}

func (server *Server) createUser(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var u models.User
	err := decoder.Decode(&u)

	if err != nil {
		// send bad response
	}
	//todo: this should return something to indicate good response
	server.userService.CreateUser(u.Email, u.Password)
}

// Handler :
func (server *Server) handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/", server.createUser)
	return mux
}

// NewServer :
func NewServer(config *models.Config, userService *service.UserService) *Server {
	return &Server{
		config:      config,
		userService: userService,
	}
}

// Start :
func (server *Server) Start() {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: server.handler(),
	}
	fmt.Println("Starting server.. on port :8080")
	httpServer.ListenAndServe()
}
