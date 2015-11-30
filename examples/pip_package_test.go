package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/matchers/pkg"
	"github.com/milosgajdos83/servpeek/resource"
)

func TestPipPackage(t *testing.T) {
	testPkg, err := resource.NewSwPkg("pip", "setuptools", "3.3")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
