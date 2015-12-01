// build linux

package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/process"
)

func TestProcess(t *testing.T) {
	if err := process.IsRunningCmd("docker"); err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := process.IsRunningCmdWithUID("docker", "root"); err != nil {
		t.Errorf("Error: %s", err)
	}
}
