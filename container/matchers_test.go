package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerImg(t *testing.T) {
	di := &DockerImg{
		Repo: "some/image",
		Tag:  "23",
	}

	assert.Equal(t, "[Docker Image] Repo: "+di.Repo+", Tag: "+di.Tag, di.String())
}

func TestDockerContainer(t *testing.T) {
	dc := &DockerContainer{
		ID:   "2432df023",
		Name: "linus_foobar",
	}

	assert.Equal(t, "[Docker Container] ID: "+dc.ID+", Name: "+dc.Name, dc.String())
}
