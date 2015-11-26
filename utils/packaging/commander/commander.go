package commander

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// PkgCommander provides interface to software package commands
type PkgCommander interface {
	// ListPkgs runs a command which queries installed packages
	// It returns command.Outer interface that can be used to parse output
	ListPkgs() command.Outer
	// QueryPkg runs a command which queries a package for various package properties
	// It returns command.Outer interface that can be used to parse output
	QueryPkg(string) command.Outer
}

// NewPkgCommander returns package manager Commander
func NewPkgCommander(pkgType string) (PkgCommander, error) {
	switch pkgType {
	case "apt", "dpkg":
		return NewAptCommander(), nil
	case "apk":
		return NewApkCommander(), nil
	case "rpm", "yum":
		return NewYumCommander(), nil
	case "pip":
		return NewPipCommander(), nil
	case "gem":
		return NewGemCommander(), nil
	}
	return nil, fmt.Errorf("PkgCommander error. Unsupported package type: %s", pkgType)
}

// BaseCommander provides basic package manager commands
type BaseCommander struct {
	// ListPkgs provides a command to list packages
	ListPkgsCmd *command.Command
	// QueryPkg provides a command to query a package
	QueryPkgCmd *command.Command
}

// ListPkgs runs a command that list installed package
// It returns command.Outer that can be used to parse the command output
func (bc *BaseCommander) ListPkgs() command.Outer {
	return bc.ListPkgsCmd.Run()
}

// QueryPkg runs a command that queries an installed package properties
// It returns command.Outer that can be used to parse the command output
func (bc *BaseCommander) QueryPkg(pkgName string) command.Outer {
	bc.QueryPkgCmd.Args = append(bc.QueryPkgCmd.Args, pkgName)
	return bc.QueryPkgCmd.Run()
}
