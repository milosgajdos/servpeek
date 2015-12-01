package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPkgCommander(t *testing.T) {
	assert := assert.New(t)
	cmderTypes := []string{"apt", "yum", "apk", "pip", "gem"}
	for _, cmderType := range cmderTypes {
		_, err := NewCommander(cmderType)
		assert.NoError(err)
	}
	// Unsupported pkg commander
	_, err := NewCommander("random")
	assert.Error(err)
}
