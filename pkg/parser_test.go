package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCmdOutParser(t *testing.T) {
	assert := assert.New(t)
	mgrTypes := []string{"apt", "yum", "apk", "pip", "gem"}
	for _, mgrType := range mgrTypes {
		c, err := NewCmdOutParser(mgrType)
		assert.NoError(err)
		assert.NotNil(c)
	}
	// Unsupported package manager
	c, err := NewCmdOutParser("random")
	assert.Error(err)
	assert.Nil(c)
}
