package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/service"
)

func TestService(t *testing.T) {
	s, err := service.NewOsSvc("docker", "upstart")
	if err != nil {
		t.Errorf("Error: %s", err
	}

	if err := s.IsRunning(s); err != nil {
		t.Errorf("Error: %s", err)
	}
}
