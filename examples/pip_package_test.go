package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/pkg"
)

func TestPipPackage(t *testing.T) {
	testPkg, err := pkg.NewSwPkg("setuptools", "3.3", "pip")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
