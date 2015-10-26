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

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
