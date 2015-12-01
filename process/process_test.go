package process

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOsProcess(t *testing.T) {
	assert := assert.New(t)
	pid := 100
	cmd := "mycmd"

	_, err := NewOsProcess(0, "mycmd")
	assert.Error(err)

	_, err = NewOsProcess(pid, "")
	assert.Error(err)

	p, err := NewOsProcess(pid, cmd)
	assert.NoError(err)
	assert.Equal(p.PID(), pid)
	assert.Equal(p.Cmd(), cmd)

	expected := fmt.Sprintf("[Process] PID: %d, Cmd: %s", pid, cmd)
	assert.Equal(expected, p.String())
}
