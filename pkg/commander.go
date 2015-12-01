package pkg

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// Commander provides interface to software package commands
type Commander interface {
	// ListPkgs runs a command which queries installed packages
	// It returns command.Outer interface that can be used to parse output
	ListPkgsOut() command.Outer
	// QueryPkgOut runs a command which queries a package for various package properties
	// It returns command.Outer interface that can be used to parse output
	QueryPkgOut(string) command.Outer
}

// NewCommander returns package manager Commander
func NewCommander(pkgType string) (Commander, error) {
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
	return nil, fmt.Errorf("Commander error. Unsupported package type: %s", pkgType)
}

// BaseCommander provides basic package manager commands
type baseCommander struct {
	// ListPkgsCmd provides a command to list packages
	ListPkgsCmd command.Commander
	// QueryPkgCmd provides a command to query a package
	QueryPkgCmd command.Commander
}

// ListPkgsOut runs a command that list installed package
// It returns command.Outer that can be used to parse the command output
func (b *baseCommander) ListPkgsOut() command.Outer {
	return b.ListPkgsCmd.Run()
}

// QueryPkgOut runs a command that queries an installed package properties
// It returns command.Outer that can be used to parse the command output
func (b *baseCommander) QueryPkgOut(pkgName string) command.Outer {
	b.QueryPkgCmd.AppendArgs(pkgName)
	return b.QueryPkgCmd.Run()
}
