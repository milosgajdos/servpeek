package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// PkgCommander provides interface to software package commands
type PkgCommander interface {
	// ListPkgs runs a command which queries installed packages
	// It returns command.Outer interface that can be used to parse output
	ListPkgsOut() command.Outer
	// QueryPkgOut runs a command which queries a package for various package properties
	// It returns command.Outer interface that can be used to parse output
	QueryPkgOut(string) command.Outer
}

// NewPkgCommander returns package manager Commander
func NewPkgCommander(pkgType string) (PkgCommander, error) {
	switch pkgType {
	case "apt":
		return NewAptCommander(), nil
	case "apk":
		return NewApkCommander(), nil
	case "yum":
		return NewYumCommander(), nil
	case "pip":
		return NewPipCommander(), nil
	case "gem":
		return NewGemCommander(), nil
	}
	return nil, fmt.Errorf("PkgCommander error. Unsupported package type: %s", pkgType)
}

// BasePkgCommander provides basic package manager commands
type basePkgCommander struct {
	// ListPkgsCmd provides a command to list packages
	ListPkgsCmd command.Commander
	// QueryPkgCmd provides a command to query a package
	QueryPkgCmd command.Commander
}

// ListPkgsOut runs a command that list installed package
// It returns command.Outer that can be used to parse the command output
func (b *basePkgCommander) ListPkgsOut() command.Outer {
	return b.ListPkgsCmd.Run()
}

// QueryPkgOut runs a command that queries an installed package properties
// It returns command.Outer that can be used to parse the command output
func (b *basePkgCommander) QueryPkgOut(pkgName string) command.Outer {
	b.QueryPkgCmd.AppendArgs(pkgName)
	return b.QueryPkgCmd.Run()
}
