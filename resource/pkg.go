package resource

import "fmt"

// Pkg interface defines API to software package
type Pkg interface {
	// Type returns package manager type
	Type() string
	// Name returns package name
	Name() string
	// Version returns package version
	Version() string
}

// SwPkg is a software package which has a name, version
// and is managed via particular package manager type
// SwPkg implements Pkg interface
type SwPkg struct {
	// package type
	pkgType string
	// package name
	name string
	// package version
	version string
}

// NewSwPkg returns SwPkg or error if unsupported package type is supplied as paramter
func NewSwPkg(pkgType, name, version string) (*SwPkg, error) {
	ptypes := []string{"apt", "apk", "dpkg", "yum", "rpm", "pip", "gem"}
	supported := make(map[string]bool)
	for _, ptype := range ptypes {
		supported[ptype] = true
	}

	if !supported[pkgType] {
		return nil, fmt.Errorf("Unsupported package type: %s", pkgType)
	}

	return &SwPkg{
		pkgType: pkgType,
		name:    name,
		version: version,
	}, nil
}

// Type returns package manager type
func (sp *SwPkg) Type() string {
	return sp.pkgType
}

// Name returns package name as returned by package manager
func (sp *SwPkg) Name() string {
	return sp.name
}

// Version returns package version
func (sp *SwPkg) Version() string {
	return sp.version
}

// String implements Stringer interface
func (sp *SwPkg) String() string {
	return fmt.Sprintf("[Package] Type: %s Name: %s Version: %s", sp.pkgType, sp.name, sp.version)
}
