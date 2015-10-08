package commander

import "fmt"

// Commander defines package manager commands
type Commander struct {
	// List all Installed packages
	ListPkgs *Command
	// QueryPkg queries info about package
	QueryPkg *Command
}

// NewCommander returns package manager Commander
func NewCommander(pkgType string) (*Commander, error) {
	switch pkgType {
	case "apt", "dpkg":
		return AptCommander()
	case "rpm", "yum":
		return YumCommander()
	case "pip":
		return PipCommander()
	case "gem":
		return GemCommander()
	}
	return nil, fmt.Errorf("Unsupported package type")
}
