package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/pkg"
)

func TestApkPackage(t *testing.T) {
	testPkg := resource.Pkg{
		Name:    "alpine-base",
		Version: "3.2.3-r0",
		Type:    "apk",
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
