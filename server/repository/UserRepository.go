package repository

import (
	"TaskBoard/server/models"
	"TaskBoard/server/util"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository :
// Layer responsible for communication
// with postgres
type UserRepository struct {
	//database *sql.DB
	database *gorm.DB
}

// NewUserRepository :
// Create a new UserRepository instance
func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{database}
}

func generateHashPassword(password string) string {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashPassword)
}

// CreateUser : Create a user with an email and password
// todo: should we return a UID?
func (userRepo *UserRepository) CreateUser(email string, password string) map[string]interface{} {
	fmt.Println("#SQL - INSERT#") //todo: Change this to some better logging

	uid, err := GenerateUID()
	if err != nil {
		fmt.Println("Error creating UID")
		return util.Message(http.StatusInternalServerError, "Error creating user")
	}
	hash := generateHashPassword(password)

	user := &models.User{
		Email:    email,
		Password: hash,
		UID:      uid,
	}

	userRepo.database.Create(user)

	if user.ID <= 0 {
		return util.Message(http.StatusInternalServerError, "Error creating user")
	}
	user.Password = ""
	user.UID = ""
	user.Token = userRepo.createToken(uid)

	gr := util.Message(http.StatusOK, "Successfully created user")
	gr["user"] = user
	return gr
}

func (userRepo *UserRepository) createToken(uid string) string {
	tkModel := models.Token{UID: uid}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tkModel)
	tokenString, _ := token.SignedString([]byte("password"))
	return tokenString
}

func (userRepo *UserRepository) validate() {

}

// Authenticate : Logs a user in if the credentials are correct
func (userRepo *UserRepository) Authenticate(email string, password string) (map[string]interface{}, error) {
	return nil, nil
}
