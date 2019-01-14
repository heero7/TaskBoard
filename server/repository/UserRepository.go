package repository

import (
	"database/sql"
	"fmt"
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

// CreateUser : Create a user with an email and password
// todo: should we return a UID?
func (userRepo *UserRepository) CreateUser(email string, password string) {
	fmt.Println("#SQL - INSERT#")
	sqlStatement := `
		INSERT INTO "Users" (email, password, uid)
		VALUES ($1, $2, $3)
	`
	uid := GenerateUID()
	userRepo.database.QueryRow(sqlStatement, email, password, uid)
	// todo: error check here
}
