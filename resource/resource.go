// package resource defines various resources
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
	return fmt.Sprintf("[Package] Name: %s, Version: %s, Type: %s", p.Name, p.Version, p.Type)
}

// File is an operating system file
type File struct {
	// Path to the physical file
	Path string
}

// Implement Stringer interface
func (f *File) String() string {
	return fmt.Sprintf("[File] Path: %s", f.Path)
}
