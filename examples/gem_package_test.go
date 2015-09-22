package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource/pkg"
)

func Test_Gem_Package(t *testing.T) {
	testPkg := &pkg.Pkg{
		Name:    "bundler",
		Version: "1.10.6",
		Type:    "gem",
	}

	if ok, err := testPkg.IsInstalled(); !ok {
		t.Errorf("%s package %s not installed: %s", testPkg.Type, testPkg.Name, err)
	}
}
