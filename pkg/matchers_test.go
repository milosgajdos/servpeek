package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInstalled(t *testing.T) {
	assert := assert.New(t)
	pkgs := []struct {
		pkgType string
		name    string
		version string
	}{
		{"apt", "cron", "3.0pl1-124ubuntu2"},
		{"yum", "grep", "2.20"},
		{"apk", "musl-utils", "1.1.11-r2"},
		{"pip", "setuptools", "3.3"},
		{"gem", "bundler", "1.10.6"},
	}

	for _, pkg := range pkgs {
		p, err := newMockPkg(pkg.pkgType, pkg.name, "query", pkg.version)
		assert.NoError(err)
		assert.NoError(IsInstalled(p))
	}

	// check package that is not installed
	p, err := newMockPkg("gem", "randpkg", "bogus", "")
	assert.NoError(err)
	assert.Error(IsInstalled(p))
}

func TestListPkgs(t *testing.T) {
	assert := assert.New(t)
	pkgMgr, err := newMockPkgManager("gem", "list")
	assert.NoError(err)

	// successfully parse output from ListPkgs
	pkgs, err := pkgMgr.ListPkgs()
	assert.NoError(err)
	assert.NotEmpty(pkgs)

	// can't parse output from ListPkgs
	pkgMgr, err = newMockPkgManager("gem", "bogus")
	assert.NoError(err)

	pkgs, err = pkgMgr.ListPkgs()
	assert.NoError(err)
	assert.Empty(pkgs)
}
