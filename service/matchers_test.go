package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRunning(t *testing.T) {
	assert := assert.New(t)

	svcs := []struct {
		name        string
		sysInitType string
	}{
		{"docker", "upstart"},
		{"docker", "systemd"},
	}

	svcStatuses := []string{"running", "stopped"}
	for _, svcStatus := range svcStatuses {
		for _, svc := range svcs {
			s, err := newMockSvc(svcStatus, svc.sysInitType, svc.name)
			assert.NoError(err)
			switch svcStatus {
			case "running":
				assert.NoError(IsRunning(s))
			case "stopped":
				assert.Error(IsRunning(s))
			}
		}
	}
}
