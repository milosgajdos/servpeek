package commander

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// PkgCommander provides package commander commands
// It defines a list of commands provided by package commanders
type PkgCommander struct {
	// List all installed packages
	ListPkgs *command.Command
	// QueryPkg queries info about a package
	QueryPkg *command.Command
}

// NewPkgCommander returns package manager Commander
func NewPkgCommander(pkgType string) (*PkgCommander, error) {
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
	return nil, fmt.Errorf("Unsupported package type: %s", pkgType)
}
