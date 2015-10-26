// package pkg provides function that check various package aspect
package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/packaging/manager"
	"github.com/milosgajdos83/servpeek/utils/packaging/parser"
)

// IsInstalled return true if all the supplied packages are installed
func IsInstalled(pkgs ...resource.Pkg) error {
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

		if len(inPkgs) > 0 {
			return nil
		}

	}
	return fmt.Errorf("Error looking up installed pacakges")
}

// IsInstalledVersion returns true if all the supplied packages
// are installed with the required version
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

		for _, inPkg := range inPkgs {
			if inPkg.Version == p.Version {
				return nil
			}
		}
	}
	return fmt.Errorf("Error looking up package version")
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
