package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/pkg"
)

func TestGemPackage(t *testing.T) {
	testPkg, err := pkg.NewSwPkg("gem", "bundler", "1.10.6")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
