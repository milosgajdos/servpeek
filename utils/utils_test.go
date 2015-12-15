package utils

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"

	"github.com/milosgajdos83/servpeek/utils/command"
	"github.com/stretchr/testify/assert"
)

func TestBuildCmd(t *testing.T) {
	cmd := "foo"
	args := []string{"bar"}
	actCmd := BuildCmd(cmd, args...)
	expCmd := &command.Command{
		Cmd:  cmd,
		Args: args,
	}
	assert.Exactly(t, expCmd, actCmd)
}

func TestRoleToID(t *testing.T) {
	assert := assert.New(t)

	name := "root"
	uid, err := RoleToID("user", name)

	switch os := runtime.GOOS; os {
	case "linux", "darwin":
		assert.NoError(err)
		assert.EqualValues(0, uid)
	case "windows":
		expErr := fmt.Sprintf("RoleToID not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
		assert.Equal(expErr, err.Error())
	}
}

func TestHashSum(t *testing.T) {
	assert := assert.New(t)
	data := `line1
	line2`

	testHashes := []struct {
		hashType string
		expected string
	}{
		{"md5", "3dfc125676228ddbac790f3b6d8d58be"},
		{"sha256", "2f6928c43c919915d452b6f2b90f7cf6640a7773c83412bf3d8ea1abfc699020"},
	}

	for _, th := range testHashes {
		r := bytes.NewBuffer([]byte(data))
		actual, err := HashSum(th.hashType, r)
		assert.NoError(err)
		assert.EqualValues(th.expected, actual)
	}
}
