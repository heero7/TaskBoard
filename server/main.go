package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 3003
	user = "admin"
	// add password here
	dbname = "TaskBoard1"
)

func main() {

	fmt.Println("Connecting to database..")
	postgresqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", postgresqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// err = db.Ping()

	// NOTES
	/*
		CASING IS EXTREMELY IMPORTANT
		If your table name is cased. i.e. like below 'Users'
		you'll need to put this in quotes.
	*/
	sqlStatement := `
		INSERT INTO "Users" (email, password, uid)
		VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, "test1@testerz.com", "abc123", "b37885dd-1e69-4c23-8755-0df68e29871b")

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully inserted!")
}
