// Package pkg provides function that check various package aspect
package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/packaging/manager"
	"github.com/milosgajdos83/servpeek/utils/packaging/parser"
)

// AllVersions is a char that specifies that the version in the package
// doesn't really matter.
const AllVersions = "*"

// IsInstalled returns nil if the package was found. Otherwise it will return
// an error.
func IsInstalled(pkgs ...resource.Pkg) error {
	for _, p := range pkgs {
		p.Version = AllVersions
	}
	return IsInstalledVersion(pkgs...)
}

// IsInstalledVersion returns nil if the versions of the given package were
// found, otherwise it will return an error.
// Note: the version "*" is special and means that doesn't really matter.
func IsInstalledVersion(pkgs ...resource.Pkg) error {
	for _, p := range pkgs {
		pkgMgr, err := manager.NewPkgManager(p.Type)
		if err != nil {
			return err
		}

		cmdParser, err := parser.NewParser(p.Type)
		if err != nil {
			return err
		}

		queryOut := pkgMgr.QueryPkg(p.Name)
		inPkgs, err := cmdParser.ParseQuery(queryOut)
		if err != nil {
			return err
		}

		if len(inPkgs) == 0 {
			return fmt.Errorf("Error looking up for %s", p.Name)
		}

		if p.Version == "*" {
			continue
		}

		for _, inPkg := range inPkgs {
			if inPkg.Version == p.Version {
				continue
			}
		}

		return fmt.Errorf("Error looking up for %s, version: %s", p.Name, p.Version)
	}
	return nil
}

// ListInstalled lists all installed packages or returns error
func ListInstalled(pkgType string) ([]*resource.Pkg, error) {
	pkgMgr, err := manager.NewPkgManager(pkgType)
	if err != nil {
		return nil, err
	}

	cmdParser, err := parser.NewParser(pkgType)
	if err != nil {
		return nil, err
	}

	listOut := pkgMgr.ListPkgs()
	pkgs, err := cmdParser.ParseList(listOut)
	if err != nil {
		return nil, err
	}

	return pkgs, nil
}
