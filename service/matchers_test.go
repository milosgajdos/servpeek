package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRunning(t *testing.T) {
	assert := assert.New(t)

	stoppedMock := &mockSysInit{status: Stopped, err: errors.New("tst")}
	svcStopped := &mockService{"mysvc", stoppedMock}
	assert.Error(IsRunning(svcStopped))

	runningMock := &mockSysInit{status: Running, err: nil}
	svcRunning := &mockService{"mysvc", runningMock}
	assert.NoError(IsRunning(svcRunning))
}
