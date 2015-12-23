package service

import (
	"io/ioutil"
	"os"
	"path"
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

	statuses := []string{"running", "stopped"}
	for _, status := range statuses {
		for _, svc := range svcs {
			currentDir, err := os.Getwd()
			assert.NoError(err)
			fixturesPath := path.Join(currentDir, "test-fixtures",
				svc.sysInitType+"-"+status+".out")
			cmdOut, err := ioutil.ReadFile(fixturesPath)
			assert.NoError(err)
			s := &mockSvc{
				name: svc.name,
				sysInit: &mockSysInit{
					sysInitType: svc.sysInitType,
					mockCommander: &mockCommander{
						StatusCmd: &mockCmd{
							out: string(cmdOut),
						},
					},
				},
			}
			switch status {
			case "running":
				assert.NoError(IsRunning(s))
			case "stopped":
				assert.Error(IsRunning(s))
			}
		}
	}
}
