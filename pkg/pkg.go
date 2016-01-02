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

// SwPkg is a software package which has a name, version
// and is managed via package manager.
// SwPkg implements Pkg interface
type SwPkg struct {
	// package manager
	manager Manager
	// package name
	name string
	// package version
	versions []string
}

// NewSwPkg returns *SwPkg. It returns error if either the requested package type is unsupported
// or package name passed as parameter is empty string. If version is empty string, it is
// ignored by matchers
func NewSwPkg(pkgType, pkgName string, pkgVersions ...string) (*SwPkg, error) {
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

	return &SwPkg{
		manager:  manager,
		name:     pkgName,
		versions: pkgVersions,
	}, nil
}

// Name returns package name as returned by package manager
func (s *SwPkg) Name() string {
	return s.name
}

// Version returns a slice of all package versions
// If the returned slice is nil, no version has been specified
func (s *SwPkg) Versions() []string {
	return s.versions
}

// Manager returns package manager that manages this type of package
func (s *SwPkg) Manager() Manager {
	return s.manager
}

// String implements Stringer interface
func (s *SwPkg) String() string {
	return fmt.Sprintf("[SwPkg] Type: %s Name: %s Version: %v",
		s.manager.Type(), s.name, s.versions)
}
