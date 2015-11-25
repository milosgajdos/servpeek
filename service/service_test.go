package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	assert := assert.New(t)

	s, err := NewOsSvc("mysvc", "upstart")
	assert.NoError(err)

	expected := "[OsSvc] Name: " + s.Name() + ", SysInit: " + s.SysInit()
	assert.Equal(expected, s.String())
}
