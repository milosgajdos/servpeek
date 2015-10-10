package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/pkg"
)

func Test_Gem_Package(t *testing.T) {
	testPkg := resource.Pkg{
		Name:    "bundler",
		Version: "1.10.6",
		Type:    "gem",
	}

	if ok, err := pkg.IsInstalled(testPkg); !ok {
		t.Errorf("%s not installed: %s", testPkg, err)
	}
}
