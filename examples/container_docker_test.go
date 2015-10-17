package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/container"
)

func Test_Docker(t *testing.T) {
	testImg := &resource.DockerImg{
		Repo: "busybox",
	}

	testContainer := &resource.DockerContainer{
		Name: "pensive_mahavira",
	}

	if ok, err := container.IsDockerImgPresent(testImg); !ok {
		t.Errorf("%s", err)
	}

	if ok, err := container.IsDockerContainerPresent(testContainer); !ok {
		t.Errorf("%s", err)
	}
}
