package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSysInit(t *testing.T) {
	assert := assert.New(t)

	initTypes := []string{"upstart", "systemd", "sysv"}
	for _, initType := range initTypes {
		s, err := NewSysInit(initType)
		assert.NoError(err)
		assert.Equal(s.Type(), initType)
	}

	_, err := NewSysInit("random")
	assert.Error(err)
}
