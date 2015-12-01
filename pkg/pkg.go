package pkg

import "fmt"

// Pkg interface defines API to software package
type Pkg interface {
	// Manager returns package manager
	Manager() Manager
	// Name returns package name
	Name() string
	// Version returns package version
	Version() string
}

// SwPkg is a software package which has a name, version
// and is managed via particular package manager type
// SwPkg implements Pkg interface
type SwPkg struct {
	// package manager
	m Manager
	// package name
	name string
	// package version
	version string
}

// NewSwPkg returns SwPkg or error if unsupported package type is supplied as paramter
func NewSwPkg(pkgType, name, version string) (*SwPkg, error) {
	ptypes := []string{"apt", "apk", "yum", "pip", "gem"}
	supported := make(map[string]bool)
	for _, ptype := range ptypes {
		supported[ptype] = true
	}

	if !supported[pkgType] {
		return nil, fmt.Errorf("Unsupported package type: %s", pkgType)
	}

	if name == "" {
		return nil, fmt.Errorf("Package name can not be empty")
	}

	manager, err := NewManager(pkgType)
	if err != nil {
		return nil, err
	}

	return &SwPkg{
		m:       manager,
		name:    name,
		version: version,
	}, nil
}

// Manager returns package manager that maintains this type of package
func (s *SwPkg) Manager() Manager {
	return s.m
}

// Name returns package name as returned by package manager
func (s *SwPkg) Name() string {
	return s.name
}

// Version returns package version
func (s *SwPkg) Version() string {
	return s.version
}

// String implements Stringer interface
func (s *SwPkg) String() string {
	return fmt.Sprintf("[Pkg] Type: %s Name: %s Version: %s", s.m.Type(), s.name, s.version)
}
