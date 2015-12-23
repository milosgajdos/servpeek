package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSysInit(t *testing.T) {
	assert := assert.New(t)

	sysInitTypes := []string{"sysv", "upstart", "systemd"}
	for _, sysInitType := range sysInitTypes {
		s, err := NewSysInit(sysInitType)
		assert.NoError(err)
		assert.Equal(s.Type(), sysInitType)
	}

	// Unsupported init types
	_, err := NewSysInit("random")
	assert.Error(err)
}
