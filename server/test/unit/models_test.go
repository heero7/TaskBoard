package unit

import (
	"TaskBoard/server/models"
	"testing"
)

func TestConfigIsNotNill(t *testing.T) {
	c := models.NewConfig()

	if c == nil {
		t.Error("Config value was nil, expected to not be nil")
	}
}
