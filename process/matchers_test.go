// +build linux

package process

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	assert := assert.New(t)

	p := &Process{
		Pid: 100,
		Cmd: "testprocess",
	}

	expected := fmt.Sprintf("[Process] PID: %d, Cmd: %s", p.Pid, p.Cmd)
	assert.Equal(expected, p.String())
}
