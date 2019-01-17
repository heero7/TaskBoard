package repository

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository :
// Layer responsible for communication
// with postgres
type UserRepository struct {
	database *sql.DB
}

// NewUserRepository :
// Create a new UserRepository instance
func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database}
}

func generateHashPassword(password string) string {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashPassword)
}

// CreateUser : Create a user with an email and password
// todo: should we return a UID?
func (userRepo *UserRepository) CreateUser(email string, password string) error {
	fmt.Println("#SQL - INSERT#") //todo: Change this to some better logging

	sqlStatement := `
		INSERT INTO "Users" (email, password, uid)
		VALUES ($1, $2, $3)
	`
	uid, err := GenerateUID()
	if err != nil {
		return err
	}
	hash := generateHashPassword(password)
	_, err = userRepo.database.Exec(sqlStatement, email, hash, uid)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (userRepo *UserRepository) Authenticate(email string, password string) {

}
