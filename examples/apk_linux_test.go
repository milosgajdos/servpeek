package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/pkg"
)

func Test_Package(t *testing.T) {
	testPkg := resource.Pkg{
		Name:    "alpine-base",
		Version: "3.2.3-r0",
		Type:    "apk",
	}

	if ok, err := pkg.IsInstalled(testPkg); !ok {
		t.Errorf("%s not installed: %s", testPkg, err)
	}
}
