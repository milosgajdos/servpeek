package main

import "github.com/milosgajdos83/servpeek/pkg"

// CheckPackages checks various properties of different kid of software pacakges
// It returns error if any of the checked properties could not be satisfied.
func CheckPackages() error {
	tstPkgs := []struct {
		pkgType string
		name    string
		version string
	}{
		{"apk", "alpine-base", "3.2.3-r0"},
		{"gem", "bundler", "1.10.6"},
		{"apt", "docker-engine", "1.8.2-0~trusty"},
		{"yum", "grep", "2.20"},
		{"pip", "setuptools", "3.3"},
	}

	for _, tstPkg := range tstPkgs {
		p, err := pkg.NewSwPkg(tstPkg.pkgType, tstPkg.name, tstPkg.version)
		if err != nil {
			return err
		}
		if err := pkg.IsInstalled(p); err != nil {
			return err
		}
	}

	return nil
}
