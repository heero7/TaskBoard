package repository

import (
	"TaskBoard/server/models"
	"TaskBoard/server/util"
	"net/http"
	"strings"

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
	//fmt.Println("#SQL - INSERT#") //todo: Change this to some better logging

	uid, err := GenerateUID()
	if err != nil {
		return util.Message(http.StatusInternalServerError, "Error creating user")
	}
	hash := generateHashPassword(password)

	user := &models.User{
		Email:    email,
		Password: hash,
		UID:      uid,
	}

	valid := userRepo.validate(user)

	if valid != nil {
		return valid
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

func (userRepo *UserRepository) validate(user *models.User) map[string]interface{} {
	if !strings.Contains(user.Email, "@") {
		return util.Message(http.StatusBadRequest, "Invalid email")
	}

	if len(user.Password) < 6 {
		return util.Message(http.StatusBadRequest, "Password length requirement not met")
	}

	// check for duplicate email
	check := &models.User{}
	err := userRepo.database.Table("users").Where("email = ?", user.Email).First(check).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return util.Message(http.StatusInternalServerError, err.Error())
	}

	if check.Email != "" {
		return util.Message(http.StatusBadRequest, "This email is already taken")
	}
	return nil
}

// Authenticate : Logs a user in if the credentials are correct
func (userRepo *UserRepository) Authenticate(email, password string) map[string]interface{} {
	user := &models.User{}
	err := userRepo.database.Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return util.Message(http.StatusNotFound, "Email address not found")
		}
		return util.Message(http.StatusInternalServerError, "Error finding email address, please try again")
	}

	// Found the email now lets compare the hash to the password receive
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		// Incorrect credentials!
		return util.Message(http.StatusInternalServerError, "Invalid credentials logged in")
	}

	// Success matching the two passwords
	user.Token = userRepo.createToken(user.UID)
	user.UID = ""
	user.Password = ""

	res := util.Message(http.StatusOK, "User was successfully logged in")
	res["user"] = user
	return res
}
