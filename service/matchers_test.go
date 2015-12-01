package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSvc(t *testing.T) {
	assert := assert.New(t)

	_, err := NewOsSvc("servicetst", "systemd")
	assert.NoError(err)
}
