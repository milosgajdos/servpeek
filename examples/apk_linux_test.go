// build linux

package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/pkg"
	"github.com/stretchr/testify/assert"
)

func TestApkPackage(t *testing.T) {
	assert := assert.New(t)
	testPkg, err := pkg.NewSwPkg("apk", "alpine-base", "3.2.3-r0")
	assert.NoError(err)

	err = pkg.IsInstalled(testPkg)
	asert.NoError(err)
}
