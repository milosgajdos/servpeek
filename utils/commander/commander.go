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
		return NewAptCommander(), nil
	case "rpm", "yum":
		return NewYumCommander(), nil
	case "pip":
		return NewPipCommander(), nil
	case "gem":
		return NewGemCommander(), nil
	}
	return nil, fmt.Errorf("Unsupported package type")
}
