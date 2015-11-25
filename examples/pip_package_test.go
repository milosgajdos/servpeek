package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/pkg"
	"github.com/stretchr/testify/assert"
)

func TestPipPackage(t *testing.T) {
	testPkg, err := pkg.NewSwPkg("pip", "setuptools", "3.3")
	assert.NoError(t, err)

	err = pkg.IsInstalled(testPkg)
	assert.NoError(t, err)
}
