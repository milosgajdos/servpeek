// Package pkg provides function that check various software package properties
package pkg

import "fmt"

// IsInstalled checks if all the packages passed in as parameters are installed
// with the required version. If pkg.Version() returns empty string, version check is ignored
// IsInstalled returns error if at least one supplied package is not installed
// or if it's verstion is different from the required version.
func IsInstalled(pkgs ...Pkg) error {
	for _, p := range pkgs {
		foundPkgs, err := p.Manager().QueryPkg(p.Name())
		if err != nil {
			return err
		}

		if len(foundPkgs) == 0 {
			return fmt.Errorf("Unable to look up %s: no package found", p)
		}

		if p.Version() == nil {
			continue
		}

	CheckPkgs:
		for _, foundPkg := range foundPkgs {
			for _, foundPkgVersion := range foundPkg.Version() {
				for _, pkgVersion := range p.Version() {
					if foundPkgVersion == pkgVersion {
						break CheckPkgs
					}
				}
			}
			return fmt.Errorf("Requested package versions: %v, found: %v",
				p.Version(), foundPkg.Version())
		}
	}
	return nil
}

// ListInstalled lists all installed packages.
// It returns error if either installed packages can't be listed
// or the output of the package manager could not be parsed
func ListInstalled(pkgType string) ([]Pkg, error) {
	pkgMgr, err := NewManager(pkgType)
	if err != nil {
		return nil, err
	}

	pkgs, err := pkgMgr.ListPkgs()
	if err != nil {
		return nil, err
	}
	return pkgs, nil
}
