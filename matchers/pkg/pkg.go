// Package pkg provides function that check various software package properties
package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/packaging/manager"
)

// IgnoreVersion is a wildcard that allows to ignore the version of the queried package
const IgnoreVersion = "*"

// IsInstalled checks if all the packages passed in as parameters are installed
// It returns error if at least supplied package is not installed
func IsInstalled(pkgs ...resource.Pkg) error {
	var ignoreVersionPkgs []resource.Pkg
	for _, p := range pkgs {
		p.Version = IgnoreVersion
		ignoreVersionPkgs = append(ignoreVersionPkgs, p)
	}
	return IsInstalledVersion(ignoreVersionPkgs...)
}

// IsInstalledVersion checks if all the packages passed in as parameters are installed
// with the required version. It returns error if at least one supplied package is not installed
// or it's verstion is different from the required version.
// Note: the version "*" is special and means that doesn't really matter.
func IsInstalledVersion(pkgs ...resource.Pkg) error {
	for _, p := range pkgs {
		pkgMgr, err := manager.NewPkgManager(p.Type)
		if err != nil {
			return err
		}

		inPkgs, err := pkgMgr.ListPkgs()
		if err != nil {
			return err
		}

		if len(inPkgs) == 0 {
			return fmt.Errorf("Unable to look up %s", p)
		}

		if p.Version == IgnoreVersion {
			continue
		}

		for _, inPkg := range inPkgs {
			if inPkg.Version == p.Version {
				continue
			}
		}
		return fmt.Errorf("Unable to look up %s", p)
	}
	return nil
}

// ListInstalled lists all installed packages.
// It returns error if either installed packages can't be listed
// or the output of the package manager could not be parsed
func ListInstalled(pkgType string) ([]*resource.Pkg, error) {
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
