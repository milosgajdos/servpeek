package pkg

import "github.com/milosgajdos83/servpeek/utils/pkg/manager"

// Pkg is a generic software package which has a name, version
// and is managed via a package manager
type Pkg struct {
	// Name is a package name
	Name string
	// Version is a package version
	Version string
	// Manager is a package manager
	Manager manager.Manager
}

func (p *Pkg) IsInstalled() bool {
	return p.Manager.CheckInstalled(p.Name)
}

func (p *Pkg) IsInstalledVersion() bool {
	return p.Manager.CheckInstalledVersion(p.Name, p.Version)
}
