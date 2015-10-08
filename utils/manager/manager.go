// package manager provides an implementation of a software
// package manager. Description sounds kind of silly redundant
// due to unfortunate naming convention classh in Go :-)
package manager

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/commander"
)

// PkgManager defines generic package manager interface
type PkgManager interface {
	// ListPkgs returns list pacakges command output
	ListPkgs() *commander.Out
	// QueryPkg returns query package command output
	QueryPkg(pkgName string) *commander.Out
}

// BasePkgManager is a package manager that implements
// basic package manager commander
type BasePkgManager struct {
	// commander provides package commander
	cmd *commander.Commander
}

// NewPkgManager returns PkgManager based on the package type
func NewPkgManager(pkgType string) (PkgManager, error) {
	switch pkgType {
	case "apt", "dpkg":
		return NewAptManager()
	case "rpm", "yum":
		return NewYumManager()
	case "pip":
		return NewPipManager()
	case "gem":
		return NewGemManager()
	}
	return nil, fmt.Errorf("Unsupported package type")
}

// AptManager implements Apt package manager
type AptManager struct {
	BasePkgManager
}

// NewAptManager returns PkgManager or fails with error
func NewAptManager() (PkgManager, error) {
	return &AptManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewAptCommander(),
		},
	}, nil
}

// YumManager implements Yum package manager
type YumManager struct {
	BasePkgManager
}

// NewYumManager returns PkgManager or fails with error
func NewYumManager() (PkgManager, error) {
	return &YumManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewYumCommander(),
		},
	}, nil
}

// PipManager implements pip package manager
type PipManager struct {
	BasePkgManager
}

// NewPipManager returns PkgManager or fails with error
func NewPipManager() (PkgManager, error) {
	return &PipManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewPipCommander(),
		},
	}, nil
}

// GemManager implements gem package manager
type GemManager struct {
	BasePkgManager
}

// NewGemManager returns PkgManager or fails with error
func NewGemManager() (PkgManager, error) {
	return &GemManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewGemCommander(),
		},
	}, nil
}

// ListPkgs runs a command which queries installed packages and returns
// the output of the command
func (bpm *BasePkgManager) ListPkgs() *commander.Out {
	return bpm.cmd.ListPkgs.Run()
}

// QueryPkg runs a command which queries a package and returns
// the output of the command
func (bpm *BasePkgManager) QueryPkg(pkgName string) *commander.Out {
	bpm.cmd.QueryPkg.Args = append(bpm.cmd.QueryPkg.Args, pkgName)
	return bpm.cmd.QueryPkg.Run()
}
