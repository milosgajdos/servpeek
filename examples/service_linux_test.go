package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/svc"
)

func Test_Service(t *testing.T) {
	dockerSvc := &resource.Svc{
		Name:    "docker",
		SysInit: "upstart",
	}
	if err := svc.IsRunning(dockerSvc); err != nil {
		t.Errorf("Error: %s", err)
	}
}
