package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/matchers/svc"
	"github.com/milosgajdos83/servpeek/resource"
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
