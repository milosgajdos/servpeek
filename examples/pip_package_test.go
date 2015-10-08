package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/pkg"
)

func Test_Pip_Package(t *testing.T) {
	testPkg := resource.Pkg{
		Name:    "setuptools",
		Version: "3.3",
		Type:    "pip",
	}

	if ok, err := pkg.IsInstalled(testPkg); !ok {
		t.Errorf("%s package %s not installed: %s", testPkg.Type, testPkg.Name, err)
	}
}
