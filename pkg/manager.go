package pkg

import "fmt"

// Manager provides software package manager interface
type Manager interface {
	// Type returns the type of the package the manager maintains
	Type() string
	// ListPkgs allows to list all installed packages on the system
	// It returns error if the installed packages can't be listed
	ListPkgs() ([]Pkg, error)
	// QueryPkg queries package database about particular package
	// It returns error if the requested package fails to be queried
	QueryPkg(pkgName string) ([]Pkg, error)
}

// NewManager returns Manager based on the requested package type passed in as parameter.
// It returns error if Manager could not be created or if provided package type is not supported.
func NewManager(pkgType string) (Manager, error) {
	switch pkgType {
	case "apt":
		return NewAptManager(), nil
	case "yum":
		return NewYumManager(), nil
	case "apk":
		return NewApkManager(), nil
	case "pip":
		return NewPipManager(), nil
	case "gem":
		return NewGemManager(), nil
	}
	return nil, fmt.Errorf("Unsupported package type: %s", pkgType)
}

// BaseManager provides basic package manager.
// It has a package command it uses to execute package management commands.
// BaseManager implements Manager interface.
type baseManager struct {
	// Commander provides package commander commands
	Commander
	// type of sw package this manager maintains
	pkgType string
}

// Type returns type of package this package manager interacts with.
func (b *baseManager) Type() string {
	return b.pkgType
}

// ListPkgs runs a command which queries installed packages.
func (b *baseManager) ListPkgs() ([]Pkg, error) {
	p, err := NewCmdOutParser(b.pkgType)
	if err != nil {
		return nil, err
	}
	return p.ParseListPkgsOut(b.ListPkgsOut())
}

// QueryPkg runs a command which queries package properties
func (b *baseManager) QueryPkg(pkgName string) ([]Pkg, error) {
	p, err := NewCmdOutParser(b.pkgType)
	if err != nil {
		return nil, err
	}
	return p.ParseQueryPkgOut(b.QueryPkgOut(pkgName))
}
