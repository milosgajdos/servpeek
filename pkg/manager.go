package pkg

import "fmt"

// PkgManager provides software package manager interface
type PkgManager interface {
	// Type returns the type of the package the manager maintains
	Type() string
	// ListPkgs allows to list all installed packages on the system
	// It returns error if the installed packages can't be listed
	ListPkgs() ([]Pkg, error)
	// QueryPkg queries package database about particular package
	// It returns error if the requested package fails to be queried
	QueryPkg(pkgName string) ([]Pkg, error)
}

// NewPkgManager returns PkgManager based on the requested package type passed in as parameter.
// It returns error if PkgManager could not be created or if provided package type is not supported.
func NewPkgManager(pkgType string) (PkgManager, error) {
	switch pkgType {
	case "apt":
		return NewAptManager()
	case "yum":
		return NewYumManager()
	case "apk":
		return NewApkManager()
	case "pip":
		return NewPipManager()
	case "gem":
		return NewGemManager()
	}
	return nil, fmt.Errorf("Unsupported package type: %s", pkgType)
}

// BasePkgManager provides basic package manager.
// It has a package command it uses to execute package management commands.
// BasePkgManager implements PkgManager interface.
type basePkgManager struct {
	// PkgCommander provides package commander commands
	PkgCommander
	// type of sw package this manager maintains
	pkgType string
}

// Type returns type of package this package manager interacts with.
func (b *basePkgManager) Type() string {
	return b.pkgType
}

// ListPkgs runs a command which queries installed packages.
func (b *basePkgManager) ListPkgs() ([]Pkg, error) {
	p, err := NewPkgOutParser(b.pkgType)
	if err != nil {
		return nil, err
	}
	return p.ParseListPkgsOut(b.ListPkgsOut())
}

// QueryPkg runs a command which queries package properties
func (b *basePkgManager) QueryPkg(pkgName string) ([]Pkg, error) {
	p, err := NewPkgOutParser(b.pkgType)
	if err != nil {
		return nil, err
	}
	return p.ParseQueryPkgOut(b.QueryPkgOut(pkgName))
}
