package commander

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPkgCommander(t *testing.T) {
	assert := assert.New(t)
	cmderTypes := []string{"apt", "dpkg", "yum", "rpm", "apk", "pip", "gem"}
	for _, cmderType := range cmderTypes {
		c, err := NewPkgCommander(cmderType)
		assert.NoError(err)
		assert.IsType(&BasePkgCommander{}, c)
	}
}
