package controllers

import (
	"TaskBoard/server/models"
	"TaskBoard/server/service"
	"TaskBoard/server/util"
	"fmt"
	"log"
	"net/http"
	"time"

	handlers "github.com/gorilla/handlers"
	gorilla "github.com/gorilla/mux"
)

// Server : Struct that holds the config and the services
type Server struct {
	config      *models.Config
	userService *service.UserService
	taskService *service.TaskService
}

func (server *Server) handler() *gorilla.Router {
	r := gorilla.NewRouter()

	uc := initUserController(server.userService)
	tc := initTasksController(server.taskService)

	r.Use(jwtAuthMiddleware)

	// BEGIN TEST ROUTES
	r.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		res := util.Message(http.StatusOK, "Hello, World")
		util.Respond(w, res)
	}).Methods("GET")
	// END TEST ROUTES

	r.HandleFunc("/api/v1/signup", uc.signUp).Methods("POST")
	r.HandleFunc("/api/v1/signin", uc.signIn).Methods("POST")
	// END USER ROUTES

	// List all tasks for a given user.. UID should be in the context of the JWT token
	r.HandleFunc("/api/v1/tasks", tc.createTask).Methods("POST")
	r.HandleFunc("/api/v1/tasks", func(w http.ResponseWriter, r *http.Request) {
		res := util.Message(http.StatusNoContent, "Not yet implemented...")
		util.Respond(w, res)
	}).Methods("GET")
	r.HandleFunc("/api/v1/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		res := util.Message(http.StatusNoContent, "Not yet implemented...")
		util.Respond(w, res)
	}).Methods("GET")
	// END TASK ROUTES
	return r
}

// NewServer : Create a new server instance
func NewServer(config *models.Config, userService *service.UserService, taskService *service.TaskService) *Server {
	return &Server{
		config:      config,
		userService: userService,
		taskService: taskService,
	}
}

// Start : Start listening on the port then serve
func (server *Server) Start() {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: server.handler(),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Starting server.. on port", httpServer.Addr)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(server.handler())))
}
