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
	tstSvcName := "mysvc"

	sysInitTypes := []string{"upstart", "systemd", "sysv"}
	for _, sysInitType := range sysInitTypes {
		s, err := NewSvc(tstSvcName, sysInitType)
		assert.NoError(err)
		assert.Equal(tstSvcName, s.Name())
		assert.Equal(sysInitType, s.SysInit().Type())
		expected := "[Svc] Name: " + s.Name() + ", SysInit: " + s.SysInit().Type()
		assert.Equal(expected, s.String())
	}

	// Unsupported SysInit
	_, err := NewSvc(tstSvcName, "mystart")
	assert.Error(err)

	// Service must have a name
	_, err = NewSvc("", "upstart")
	assert.Error(err)
}

func TestName(t *testing.T) {
	assert := assert.New(t)
	tstSvcName := "mysvc"
	s, err := NewSvc(tstSvcName, "upstart")
	assert.NoError(err)
	assert.Equal(tstSvcName, s.Name())
}

func TestSysInitType(t *testing.T) {
	assert := assert.New(t)
	tstSvcName := "mysvc"
	tstInit := "upstart"

	s, err := NewSvc(tstSvcName, tstInit)
	assert.NoError(err)
	assert.Equal(tstInit, s.SysInit().Type())

	// Service type can not be empty
	_, err = NewSvc(tstSvcName, "")
	assert.Error(err)
}
