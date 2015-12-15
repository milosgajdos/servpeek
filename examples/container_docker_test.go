package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/container"
)

func Test_Docker(t *testing.T) {
	testImg := &container.DockerImg{
		Repo: "busybox",
	}

	if err := container.IsDockerImgPresent(testImg); err != nil {
		t.Errorf("%s", err)
	}

	testContainer := &container.DockerContainer{
		Name: "pensive_mahavira",
	}

	if err := container.IsDockerContainerPresent(testContainer); err != nil {
		t.Errorf("%s", err)
	}
}
