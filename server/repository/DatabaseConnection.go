package repository

import (
	"TaskBoard/server/models"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// ConnectPostgresDatabase :
// Establishes connection to the Postgres database
func ConnectPostgresDatabase(config *models.Config) (*sql.DB, error) {
	fmt.Println("Getting config info...")
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	return sql.Open("postgres", dbinfo)
}

// GenerateUID :
// Generates a unique ID in form a string
func GenerateUID() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil

}
