package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/manager"
	"github.com/milosgajdos83/servpeek/utils/parser"
)

// IsInstalled return true if all the supplied packages are installed
func IsInstalled(pkgs ...resource.Pkg) (bool, error) {
	for _, p := range pkgs {
		pkgMgr, err := manager.NewPkgManager(p.Type)
		if err != nil {
			return false, err
		}

		cmdParser, err := parser.NewParser(p.Type)
		if err != nil {
			return false, err
		}

		queryOut := pkgMgr.QueryPkg(p.Name)
		inPkgs, err := cmdParser.ParseQuery(queryOut)
		if err != nil {
			return false, err
		}

		if len(inPkgs) > 0 {
			return true, nil
		}

	}
	return false, fmt.Errorf("Error looking up installed pacakges")
}

// IsInstalledVersion returns true if all the supplied packages
// are installed with the required version
func IsInstalledVersion(pkgs ...resource.Pkg) (bool, error) {
	for _, p := range pkgs {
		pkgMgr, err := manager.NewPkgManager(p.Type)
		if err != nil {
			return false, err
		}

		cmdParser, err := parser.NewParser(p.Type)
		if err != nil {
			return false, err
		}

		queryOut := pkgMgr.QueryPkg(p.Name)
		inPkgs, err := cmdParser.ParseQuery(queryOut)
		if err != nil {
			return false, err
		}

		for _, inPkg := range inPkgs {
			if inPkg.Version == p.Version {
				return true, nil
			}
		}
	}
	return false, fmt.Errorf("Error looking up package version")
}

// ListPackages lists all installed packages or returns error
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
	pkgs, err := cmdParser.ParseQuery(listOut)
	if err != nil {
		return nil, err
	}

	return pkgs, nil
}
