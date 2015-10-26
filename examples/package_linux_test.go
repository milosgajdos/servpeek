package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/pkg"
)

func Test_Package(t *testing.T) {
	testPkg := resource.Pkg{
		Name:    "docker-engine",
		Version: "1.8.2-0~trusty",
		Type:    "apt",
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
