// build linux

package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/matchers/pkg"
	"github.com/milosgajdos83/servpeek/resource"
)

func TestApkPackage(t *testing.T) {
	testPkg, err := resource.NewSwPkg("apk", "alpine-base", "3.2.3-r0")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
