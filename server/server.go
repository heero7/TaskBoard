package main

import (
	"TaskBoard/server/models"
	"TaskBoard/server/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	gorilla "github.com/gorilla/mux"
)

// Server :
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

func contentTypeMiddleware(nextMethod http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		nextMethod.ServeHTTP(w, r)
	})
}

func (server *Server) createUser(writer http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var u models.User
	err := decoder.Decode(&u)

	if err != nil {
		// send bad response
		br := createResponse(http.StatusInternalServerError, err.Error())
		writer.Write(br)
		return
	}
	//todo: this should return something to indicate good response
	server.userService.CreateUser(u.Email, u.Password)
}

// Handler :
func (server *Server) handler() *gorilla.Router {
	r := gorilla.NewRouter()
	r.Use(contentTypeMiddleware)
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(createResponse(http.StatusOK, "Hello, World"))
	}).Methods("GET")
	r.HandleFunc("/api/v1/", server.createUser).Methods("POST")
	return r
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
	// might want to just do the handler and listenAndServe here
	// Why? Need to do research and see if you can ensure Methods are POST
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: server.handler(),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting server.. on port :8080")
	log.Fatal(httpServer.ListenAndServe())
}
