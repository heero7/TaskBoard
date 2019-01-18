package util

import (
	"encoding/json"
	"net/http"
)

// Message : Returns a message back to the client
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond : Responds in a json format
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
