package models

import (
	"github.com/jinzhu/gorm"
)

// User : This is a user struct model used to
// post to the database. This is why the Password
// property exists.
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	UID      string
	Token    string `sql:"-"`
}
