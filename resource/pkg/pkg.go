package pkg

import "github.com/milosgajdos83/servpeek/utils/pkg/manager"

// Pkg is a generic software package which has a name, version
// and is managed via a package manager
type Pkg struct {
	// Name is a package name
	Name string
	// Version is a package version
	Version string
	// Type is a package type
	Type string
}

// IsInstalled return true if the package is installed
func (p *Pkg) IsInstalled() (bool, error) {
	pkgMgr, err := manager.NewPkgManager(p.Type)
	if err != nil {
		return false, err
	}
	return pkgMgr.CheckInstalled(p.Name)
}

// IsInstalledVersion returns true installed package is
// of the given version
func (p *Pkg) IsInstalledVersion() (bool, error) {
	pkgMgr, err := manager.NewPkgManager(p.Type)
	if err != nil {
		return false, err
	}
	return pkgMgr.CheckInstalledVersion(p.Name, p.Version)
}
