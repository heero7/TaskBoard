package unit

import "testing"
import u "TaskBoard/server/util"

func TestCreateMessage(t *testing.T) {
	message := "Success"
	code := 200
	m := u.Message(code, message)

	if m == nil {
		t.Error("Expected to be populated, got nil")
	}

	if m["status"] != code {
		t.Error("Expected status of 200, got", m["status"])
	}

	if m["message"] != message {
		t.Error("Expected message,", message)
		t.Error("Got,", m["message"])
	}
}
