package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/manager"
	"github.com/milosgajdos83/servpeek/utils/parser"
)

// Pkg is a generic software package which has a name, version
// and is managed via a package manager
type Pkg struct {
	// package name
	Name string
	// package version
	Version string
	// package type
	Type string
}

// IsInstalled return true if the package is Installed
func (p *Pkg) IsInstalled() (bool, error) {
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

	return false, fmt.Errorf("Package %s not found", p.Name)
}

// IsInstalledVersion returns true if the installed package
// has the required version
func (p *Pkg) IsInstalledVersion() (bool, error) {
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

	return false, fmt.Errorf("Package %s verion %s not found", p.Name, p.Version)
}

// ListPackages lists all installed packages or returns error
// It infers package type if provided pkgMgr type is empty
func ListInstalled(pkgType string) ([]*Pkg, error) {
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
