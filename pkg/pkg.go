package pkg

import "fmt"

var supportedPkgTypes = map[string]bool{
	"apt": true,
	"apk": true,
	"yum": true,
	"pip": true,
	"gem": true,
}

// Pkg interface defines simple API to software package
type Pkg interface {
	// Name returns package name
	Name() string
	// Versions returns slice of all package versions
	Versions() []string
	// Manager returns package manager
	Manager() Manager
}

// Package is a software package which has a name, version
// and is managed via package manager.
// Package implements Pkg interface
type Package struct {
	// package manager
	manager Manager
	// package name
	name string
	// package version
	versions []string
}

// NewPackage returns *Package. It returns error if either the requested package type is unsupported
// or package name passed as parameter is empty string. If version is empty string, it is
// ignored by matchers
func NewPackage(pkgType, pkgName string, pkgVersions ...string) (*Package, error) {
	if !supportedPkgTypes[pkgType] {
		return nil, fmt.Errorf("Unsupported package type: %s", pkgType)
	}

	if pkgName == "" {
		return nil, fmt.Errorf("Package name can not be empty!")
	}

	manager, err := NewManager(pkgType)
	if err != nil {
		return nil, err
	}

	return &Package{
		manager:  manager,
		name:     pkgName,
		versions: pkgVersions,
	}, nil
}

// Name returns package name as returned by package manager
func (s *Package) Name() string {
	return s.name
}

// Version returns a slice of all package versions
// If the returned slice is nil, no version has been specified
func (s *Package) Versions() []string {
	return s.versions
}

// Manager returns package manager that manages this type of package
func (s *Package) Manager() Manager {
	return s.manager
}

// String implements Stringer interface
func (s *Package) String() string {
	return fmt.Sprintf("[Package] Type: %s Name: %s Version: %v",
		s.manager.Type(), s.name, s.versions)
}
