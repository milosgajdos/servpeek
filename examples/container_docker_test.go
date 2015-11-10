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

	if err := container.IsDockerImgPresent(testImg); err != nil {
		t.Errorf("%s", err)
	}

	if err := container.IsDockerContainerPresent(testContainer); err != nil {
		t.Errorf("%s", err)
	}
}
