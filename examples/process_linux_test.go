package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/matchers/process"
)

func Test_Process(t *testing.T) {
	if err := process.IsRunningCmd("docker"); err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := process.IsRunningCmdWithUID("docker", "root"); err != nil {
		t.Errorf("Error: %s", err)
	}
}
