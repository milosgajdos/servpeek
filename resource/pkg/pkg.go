package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/pkg/manager"
)

// Pkg is a generic software package which has a name, version
// and is managed via a package manager
type Pkg struct {
	Name    string
	Version string
	Type    string
}

// IsInstalled return true if the package is installed
func (p *Pkg) IsInstalled() (bool, error) {
	pkgMgr, err := manager.NewPkgManager(p.Type)
	if err != nil {
		return false, err
	}

	pkgs, err := pkgMgr.QueryInstalled(p.Name)
	if err != nil {
		return false, err
	}

	for _, pkg := range pkgs {
		if pkg.Name == p.Name {
			return true, nil
		}
	}

	return false, fmt.Errorf("Package %s not found", p.Name)
}

// IsInstalledVersion returns true installed package has the given version
func (p *Pkg) IsInstalledVersion() (bool, error) {
	pkgMgr, err := manager.NewPkgManager(p.Type)
	if err != nil {
		return false, err
	}

	pkgs, err := pkgMgr.QueryInstalled(p.Name)
	if err != nil {
		return false, err
	}

	for _, pkg := range pkgs {
		if pkg.Version == p.Version {
			return true, nil
		}
	}

	return false, fmt.Errorf("Package %s verion %s not found", p.Name, p.Version)
}
