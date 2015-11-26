package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSvc(t *testing.T) {
	assert := assert.New(t)

	s := &Svc{
		Name:    "servicetst",
		SysInit: "systemd",
	}

	assert.Equal("[Service] Name: "+s.Name+", SysInit: "+s.SysInit, s.String())
}
