// build linux

package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/pkg"
)

func TestApkPackage(t *testing.T) {
	testPkg, err := pkg.NewSwPkg("alpine-base", "3.2.3-r0", "apk")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
