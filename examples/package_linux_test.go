// build linux

package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/pkg"
	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	testPkg, err := pkg.NewSwPkg("apt", "docker-engine", "1.8.2-0~trusty")
	assert.NoError(t, err)

	err = pkg.IsInstalled(testPkg)
	assert.NoError(t, err)
}
