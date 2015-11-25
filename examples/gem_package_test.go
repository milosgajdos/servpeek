package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGemPackage(t *testing.T) {
	testPkg, err := pkg.NewSwPkg("gem", "bundler", "1.10.6")
	assert.NoError(t, err)

	err = pkg.IsInstalled(testPkg)
	assert.NoError(t, err)
}
