// Package manager provides an implementation of a software package manager.
// Description sounds kind of silly redundant due to unfortunate naming convention clash in Go :-)
package manager

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
	"github.com/milosgajdos83/servpeek/utils/packaging/commander"
)

// PkgManager provides package manager interface
type PkgManager interface {
	// ListPkgs runs command that lists installed packages
	ListPkgs() *command.Out
	// QueryPkg runs command that queries info about installed package
	QueryPkg(pkgName string) *command.Out
}

// BasePkgManager implements basic package manager commands
type BasePkgManager struct {
	// cmd provides package commander
	cmd *commander.PkgCommander
}

// ListPkgs runs a command which queries installed packages
// It returns its output that can be parsed
func (bpm *BasePkgManager) ListPkgs() *command.Out {
	return bpm.cmd.ListPkgs.Run()
}

// QueryPkg runs a command which queries a package
// It returns the output that can be parsed
func (bpm *BasePkgManager) QueryPkg(pkgName string) *command.Out {
	bpm.cmd.QueryPkg.Args = append(bpm.cmd.QueryPkg.Args, pkgName)
	return bpm.cmd.QueryPkg.Run()
}

// NewPkgManager returns PkgManager based on the package type
// It returns error if the PkgManager could not be created or required package type is not supported
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

// aptManager implements Apt package manager
type aptManager struct {
	BasePkgManager
}

// NewAptManager returns PkgManager or fails with error
func NewAptManager() (PkgManager, error) {
	return &aptManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewAptCommander(),
		},
	}, nil
}

// yumManager implements Yum package manager
type yumManager struct {
	BasePkgManager
}

// NewYumManager returns PkgManager or fails with error
func NewYumManager() (PkgManager, error) {
	return &yumManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewYumCommander(),
		},
	}, nil
}

// pipManager implements pip package manager
type pipManager struct {
	BasePkgManager
}

// NewPipManager returns PkgManager or fails with error
func NewPipManager() (PkgManager, error) {
	return &pipManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewPipCommander(),
		},
	}, nil
}

// apkManager implements apk package manager
type apkManager struct {
	BasePkgManager
}

// NewApkManager returns PkgManager or fails with error
func NewApkManager() (PkgManager, error) {
	return &apkManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewApkCommander(),
		},
	}, nil
}

// GemManager implements gem package manager
type gemManager struct {
	BasePkgManager
}

// NewGemManager returns PkgManager or fails with error
func NewGemManager() (PkgManager, error) {
	return &gemManager{
		BasePkgManager: BasePkgManager{
			cmd: commander.NewGemCommander(),
		},
	}, nil
}
