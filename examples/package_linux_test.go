// build linux

package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/pkg"
)

func TestPackage(t *testing.T) {
	testPkg, err := pkg.NewSwPkg("docker-engine", "1.8.2-0~trusty", "apt")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := pkg.IsInstalled(testPkg); err != nil {
		t.Errorf("Error: %s", err)
	}
}
