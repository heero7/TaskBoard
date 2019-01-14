package models

// User : This is a user struct model used to
// post to the database. This is why the Password
// property exists.
type User struct {
	Email    string
	Password string
}
