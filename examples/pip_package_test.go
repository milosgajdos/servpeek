package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/matchers/pkg"
	"github.com/milosgajdos83/servpeek/resource"
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
