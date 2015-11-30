// build linux

package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/matchers/pkg"
	"github.com/milosgajdos83/servpeek/resource"
)

func TestPackage(t *testing.T) {
	testPkg, err := resource.NewSwPkg("apt", "docker-engine", "1.8.2-0~trusty")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
