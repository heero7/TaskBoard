package models

// Config : Configuration for the database needed to
// connect. This is specifically for Postgres.
type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

// NewConfig : Creates a new configuration struct
func NewConfig() *Config {
	// todo: not implemented
	return nil
}
