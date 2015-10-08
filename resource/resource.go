// package resource defines compute resources
package resource

// Pkg is a generic software package which has a name, version
// and is managed via a package manager
type Pkg struct {
	// package name
	Name string
	// package version
	Version string
	// package type
	Type string
}
