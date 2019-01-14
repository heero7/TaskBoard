package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config : Configuration for the database needed to
// connect. This is specifically for Postgres.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

// NewConfig : Creates a new configuration struct
func NewConfig() *Config {
	dbConfig, err := os.Open("./config/db.config.json")

	if err != nil {
		// error getting the file
		fmt.Println("Error opening config json,", err.Error())
	}

	defer dbConfig.Close()

	jsonParser := json.NewDecoder(dbConfig)
	c := Config{}

	if err = jsonParser.Decode(&c); err != nil {
		fmt.Println("Error parsing config json,", err.Error())
	}

	// Print our config information
	fmt.Printf("Config information... %s %s %s %s \n", c.Host, c.Port, c.User, c.Dbname)

	return &c
}
