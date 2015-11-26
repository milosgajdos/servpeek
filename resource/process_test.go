package resource

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

	str := fmt.Sprintf("[Process] PID: %d, Cmd: %s", p.Pid, p.Cmd)
	assert.Equal(str, p.String())
}
