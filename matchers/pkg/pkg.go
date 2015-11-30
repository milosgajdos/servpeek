// Package pkg provides function that check various software package properties
package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/packaging/manager"
)

// IsInstalled checks if all the packages passed in as parameters are installed
// with the required version. If pkg.Version() returns empty string, version check is ignored
// IsInstalled returns error if at least one supplied package is not installed
// or if it's verstion is different from the required version.
func IsInstalled(pkgs ...resource.Pkg) error {
	for _, p := range pkgs {
		pkgMgr, err := manager.NewPkgManager(p.Type())
		if err != nil {
			return err
		}

		inPkgs, err := pkgMgr.QueryPkg(p.Name())
		if err != nil {
			return err
		}

		if len(inPkgs) == 0 {
			return fmt.Errorf("Unable to look up %s: no package found", p)
		}

		if p.Version() == "" {
			continue
		}

		for _, inPkg := range inPkgs {
			if inPkg.Version() == p.Version() {
				continue
			}
		}
	}
	return nil
}

// ListInstalled lists all installed packages.
// It returns error if either installed packages can't be listed
// or the output of the package manager could not be parsed
func ListInstalled(pkgType string) ([]resource.Pkg, error) {
	pkgMgr, err := manager.NewPkgManager(pkgType)
	if err != nil {
		return nil, err
	}

	pkgs, err := pkgMgr.ListPkgs()
	if err != nil {
		return nil, err
	}

	return pkgs, nil
}
