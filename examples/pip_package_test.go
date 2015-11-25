package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/pkg"
)

func TestPipPackage(t *testing.T) {
	testPkg := resource.Pkg{
		Name:    "setuptools",
		Version: "3.3",
		Type:    "pip",
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
