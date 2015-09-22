package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource/pkg"
)

func Test_Package(t *testing.T) {
	testPkg := &pkg.Pkg{
		Name:    "docker-engine",
		Version: "1.8.2-0~trusty",
		Type:    "apt",
	}

	if ok, err := testPkg.IsInstalled(); !ok {
		t.Errorf("Package %s : %s not installed: %s", testPkg.Name, testPkg.Version, err)
	}
}
