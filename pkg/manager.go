package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// Manager provides software defines manager interface
type Manager interface {
	// Type returns the type of the package the manager maintains
	Type() string
	// ListPkgs allows to list all installed packages on the system
	// It returns a list of all installed packages found in package database
	// It returns error if the installed packages can't be listed
	ListPkgs() ([]Pkg, error)
	// QueryPkg queries package database about particular package
	// It returns a list of packages that match the name provided as argument
	// It returns error if the requested package fails to be queried
	QueryPkg(pkgName string) ([]Pkg, error)
}

// NewManager returns Manager based on the requested package type passed in as parameter.
// It returns error if Manager could not be created or if requested package type is not supported.
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

// BaseManager provides simple implementation of package manager.
// It has a package command it uses to execute package management commands.
// baseManager implements Manager interface.
type baseManager struct {
	// ListPkgsCmd runs command that lists installed pacakges
	ListPkgsCmd command.Command
	// QueryPkgCmd runs command that queries a package
	QueryPkgCmd command.Command
	// parser parses output of manager commands
	parser CmdOutParser
	// type of sw package this manager maintains
	pkgType string
}

// Type returns type of package this package manager manages
func (b *baseManager) Type() string {
	return b.pkgType
}

// ListPkgs runs a command which queries installed packages.
func (b *baseManager) ListPkgs() ([]Pkg, error) {
	return b.parser.ParseListPkgsOut(b.ListPkgsCmd.Run())
}

// QueryPkg runs a command which queries package properties
func (b *baseManager) QueryPkg(pkgName string) ([]Pkg, error) {
	b.QueryPkgCmd.AppendArgs(pkgName)
	return b.parser.ParseQueryPkgOut(b.QueryPkgCmd.Run())
}
