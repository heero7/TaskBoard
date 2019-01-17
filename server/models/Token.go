package models

import jwt "github.com/dgrijalva/jwt-go"

// Token : JWT struct sent to the user to authenticate
type Token struct {
	UID string
	jwt.StandardClaims
}
