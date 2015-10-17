package resource

import "fmt"

// Pkg is a software package which has a name, version
// and is managed via a package manager
type Pkg struct {
	// package name
	Name string
	// package version
	Version string
	// package type
	Type string
}

// Implement Stringer interface
func (p *Pkg) String() string {
	return fmt.Sprintf("[Package] Type: %s Name: %s Version: %s", p.Type, p.Name, p.Version)
}
