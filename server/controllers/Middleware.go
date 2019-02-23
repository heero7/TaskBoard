package controllers

import (
	"TaskBoard/server/models"
	"TaskBoard/server/util"
	"context"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type contextKey string

// todo : figure out this underlying type
// probably needs to be its own type or model
func (c contextKey) String() string {
	return string(c)
}

// todo: figure out how to implement this
func contentTypeMiddleware(nextMethod http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		nextMethod.ServeHTTP(w, r)
	})
}

func jwtAuthMiddleware(nextMethod http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doNotAuth := []string{"/", "/api/v1/signup", "/api/v1/signin"}
		requestPath := r.URL.Path

		// Do not need to authenticate for these routes
		for _, value := range doNotAuth {
			if value == requestPath {
				nextMethod.ServeHTTP(w, r)
				return
			}
		}

		tkHeader := r.Header.Get("Authorization")

		// Missing token in header
		if tkHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			br := util.Message(http.StatusForbidden, "Missing auth token!")
			util.Respond(w, br)
			return
		}

		// The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		// Not exactly why this is the case
		split := strings.Split(tkHeader, " ")
		if len(split) != 2 {
			w.WriteHeader(http.StatusForbidden)
			br := util.Message(http.StatusForbidden, "Incorrect or malformed auth token")
			util.Respond(w, br)
			return
		}

		// Only use the token part located at the second position
		tokenPt := split[1]
		tokenModel := &models.Token{}

		// secret := os.Getenv("token_password") look into setting this through config?

		token, err := jwt.ParseWithClaims(tokenPt, tokenModel, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			br := util.Message(http.StatusForbidden, "Invalid auth token")
			util.Respond(w, br)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			br := util.Message(http.StatusForbidden, "Invalid auth token")
			util.Respond(w, br)
			return
		}

		//fmt.Sprintf("User %", tokenModel.UID) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tokenModel.UID)
		r = r.WithContext(ctx)
		nextMethod.ServeHTTP(w, r)
	})
}
