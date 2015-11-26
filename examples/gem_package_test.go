package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/matchers/pkg"
	"github.com/milosgajdos83/servpeek/resource"
)

func TestGemPackage(t *testing.T) {
	testPkg := resource.Pkg{
		Name:    "bundler",
		Version: "1.10.6",
		Type:    "gem",
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
