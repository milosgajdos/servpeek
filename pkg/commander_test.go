package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPkgCommander(t *testing.T) {
	assert := assert.New(t)
	cmderTypes := []string{"apt", "yum", "apk", "pip", "gem"}
	for _, cmderType := range cmderTypes {
		_, err := NewPkgCommander(cmderType)
		assert.NoError(err)
	}
	// Unsupported pkg commander
	_, err := NewPkgCommander("random")
	assert.Error(err)
}
