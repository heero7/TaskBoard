package main

import (
	"TaskBoard/server/models"
	"TaskBoard/server/service"
	"TaskBoard/server/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	gorilla "github.com/gorilla/mux"
)

// Server : Struct that holds the config and the services
type Server struct {
	config      *models.Config
	userService *service.UserService
}

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func createResponse(status int, message string) []byte {
	res := response{Status: status, Message: message}
	js, _ := json.Marshal(&res)
	return js
}

func (server *Server) createUser(writer http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var u models.User

	err := decoder.Decode(&u)

	if err != nil && err == io.EOF {
		// send bad response
		br := util.Message(http.StatusBadRequest, "Empty request body")
		util.Respond(writer, br)
		return
	}
	//todo: this should return something to indicate good response

	res := server.userService.CreateUser(u.Email, u.Password)

	if res["status"] == 500 {
		// send bad response
		br := createResponse(http.StatusInternalServerError, "Could not create user")
		writer.Write(br)
		return
	}
	br := createResponse(http.StatusOK, fmt.Sprintf("Success creating user %s", u.Email))
	writer.Write(br)
	return
}

func signIn(writer http.ResponseWriter, request *http.Request) {

}

func (server *Server) handler() *gorilla.Router {
	r := gorilla.NewRouter()
	r.Use(jwtAuthMiddleware)

	// BEGIN TEST ROUTES
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(createResponse(http.StatusOK, "Hello, World"))
	}).Methods("GET")
	// END TEST ROUTES

	// todo: Place user routes
	r.HandleFunc("/api/v1/signup", server.createUser).Methods("POST")
	r.HandleFunc("/api/v1/signin", func(w http.ResponseWriter, r *http.Request) {
		w.Write(createResponse(http.StatusNoContent, "Not yet implemented..."))
	})
	// END USER ROUTES

	// todo: Place task routes
	// List all tasks for a given user.. UID should be in the context of the JWT token
	r.HandleFunc("/api/v1/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Write(createResponse(http.StatusNoContent, "Not yet implemented..."))
	}).Methods("GET")
	r.HandleFunc("/api/v1/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write(createResponse(http.StatusNoContent, "Not yet implemented..."))
	}).Methods("GET")
	// END TASK ROUTES
	return r
}

// NewServer : Create a new server instance
func NewServer(config *models.Config, userService *service.UserService) *Server {
	return &Server{
		config:      config,
		userService: userService,
	}
}

// Start : Start listening on the port then serve
func (server *Server) Start() {
	// might want to just do the handler and listenAndServe here
	// Why? Need to do research and see if you can ensure Methods are POST
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: server.handler(),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting server.. on port", httpServer.Addr)
	log.Fatal(httpServer.ListenAndServe())
}
