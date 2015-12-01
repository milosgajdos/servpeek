package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	assert := assert.New(t)

	statuses := []struct {
		s        Status
		expected string
	}{
		{Running, "running"},
		{Stopped, "stopped"},
		{Status(123), "unknown"},
	}

	for _, status := range statuses {
		assert.Equal(status.s.String(), status.expected)
	}
}

func TestService(t *testing.T) {
	assert := assert.New(t)
	svcName := "mysvc"

	initTypes := []string{"upstart", "systemd", "sysv"}
	for _, initType := range initTypes {
		s, err := NewOsSvc(svcName, initType)
		assert.NoError(err)
		expected := "[OsSvc] Name: " + s.Name() + ", SysInit: " + s.SysInit().Type()
		assert.Equal(expected, s.String())
	}

	_, err := NewOsSvc(svcName, "mystart")
	assert.Error(err)

	_, err = NewOsSvc("", "uptart")
	assert.Error(err)
}

func TestName(t *testing.T) {
	assert := assert.New(t)
	svcName := "mysvc"
	s, err := NewOsSvc(svcName, "upstart")
	assert.Equal(svcName, s.Name())

	_, err = NewOsSvc("", "upstart")
	assert.Error(err)
}

func TestSysInitType(t *testing.T) {
	assert := assert.New(t)
	svcName := "mysvc"
	tstInit := "upstart"

	s, err := NewOsSvc(svcName, tstInit)
	assert.NoError(err)
	assert.Equal(tstInit, s.SysInit().Type())

	_, err = NewOsSvc(svcName, "")
	assert.Error(err)
}
