package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource/svc"
)

func Test_Service(t *testing.T) {
	if err := svc.IsRunning("docker"); err != nil {
		t.Errorf("Error: %s", err)
	}
}
