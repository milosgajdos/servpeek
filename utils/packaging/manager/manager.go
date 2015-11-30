// Package manager provides functions that allow running various
// software package manager commands
package manager

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/packaging/commander"
	"github.com/milosgajdos83/servpeek/utils/packaging/parser"
)

// PkgManager provides software package manager interface
type PkgManager interface {
	// ListPkgs allows to list all installed packages on the system
	// It returns error if the installed packages can't be listed
	ListPkgs() ([]resource.Pkg, error)
	// QueryPkg queries package database about particular package
	// It returns error if the requested package fails to be queried
	QueryPkg(pkgName string) ([]resource.Pkg, error)
}

// NewPkgManager returns PkgManager based on the requested package type passed in as parameter.
// It returns error if PkgManager could not be created or if provided package type is not supported.
func NewPkgManager(pkgType string) (PkgManager, error) {
	switch pkgType {
	case "apt", "dpkg":
		return NewAptManager()
	case "rpm", "yum":
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

// BasePkgManager implements basic package manager commands
// It implements PkgManager interface
type BasePkgManager struct {
	// cmd provides package commander
	cmd commander.PkgCommander
	// parser provides package command parser
	p parser.PkgOutParser
}

// ListPkgs runs a command which queries installed packages.
func (bpm *BasePkgManager) ListPkgs() ([]resource.Pkg, error) {
	return bpm.p.ParseList(bpm.cmd.ListPkgs())
}

// QueryPkg runs a command which queries package properties
func (bpm *BasePkgManager) QueryPkg(pkgName string) ([]resource.Pkg, error) {
	return bpm.p.ParseQuery(bpm.cmd.QueryPkg(pkgName))
}

// aptManager implements Apt package manager
type aptManager struct {
	BasePkgManager
}

// NewAptManager returns apt PkgManager or fails with error
func NewAptManager() (PkgManager, error) {
	return &aptManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewAptCommander(),
			p:   parser.NewAptParser(),
		},
	}, nil
}

// yumManager implements Yum package manager
type yumManager struct {
	BasePkgManager
}

// NewYumManager returns yum PkgManager or fails with error
func NewYumManager() (PkgManager, error) {
	return &yumManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewYumCommander(),
			p:   parser.NewYumParser(),
		},
	}, nil
}

// pipManager implements pip package manager
type pipManager struct {
	BasePkgManager
}

// NewPipManager returns pip PkgManager or fails with error
func NewPipManager() (PkgManager, error) {
	return &pipManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewPipCommander(),
			p:   parser.NewPipParser(),
		},
	}, nil
}

// apkManager implements apk package manager
type apkManager struct {
	BasePkgManager
}

// NewApkManager returns apk PkgManager or fails with error
func NewApkManager() (PkgManager, error) {
	return &apkManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewApkCommander(),
			p:   parser.NewApkParser(),
		},
	}, nil
}

// GemManager implements gem package manager
type gemManager struct {
	BasePkgManager
}

// NewGemManager returns gem PkgManager or fails with error
func NewGemManager() (PkgManager, error) {
	return &gemManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewGemCommander(),
			p:   parser.NewGemParser(),
		},
	}, nil
}
