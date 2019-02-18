package controllers

import (
	"TaskBoard/server/models"
	"TaskBoard/server/util"
	"encoding/json"
	"io"
	"net/http"
)

// UserController : Handles all the routes for users
type UserController struct {
	server *Server
}

func initUserController(server *Server) *UserController {
	return &UserController{
		server: server,
	}
}

// SignIn : Route to sign in a user and send a token
func (uc *UserController) signIn(writer http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var u models.User

	err := decoder.Decode(&u)

	if err != nil && err == io.EOF {
		// send bad response
		br := util.Message(http.StatusBadRequest, "Empty request body")
		writer.WriteHeader(http.StatusInternalServerError)
		util.Respond(writer, br)
		return
	}

	res := uc.server.userService.CreateUser(u.Email, u.Password)

	if res["status"] == 500 {
		// send bad response
		br := util.Message(http.StatusInternalServerError, "Could not create user")
		writer.WriteHeader(http.StatusInternalServerError)
		util.Respond(writer, br)
		return
	}

	util.Respond(writer, res)
}

// SignUp : Route to create a user with an email and password
func (uc *UserController) signUp(writer http.ResponseWriter, request *http.Request) {
	//setupResponse(&writer, request)

	decoder := json.NewDecoder(request.Body)
	var u models.User

	err := decoder.Decode(&u)

	if err != nil && err == io.EOF {
		// send bad response
		br := util.Message(http.StatusBadRequest, "Empty request body")
		writer.WriteHeader(http.StatusInternalServerError)
		util.Respond(writer, br)
		return
	}

	res := uc.server.userService.Authenticate(u.Email, u.Password)

	if res["status"] == 404 || res["status"] == 500 {
		util.Respond(writer, res)
		return
	}
	util.Respond(writer, res)
}
