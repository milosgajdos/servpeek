package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInstalled(t *testing.T) {
	assert := assert.New(t)
	var listPkgs, queryPkgs []Pkg
	tstName := "mypkg"
	tstVersion := "0.0.1"

	pkgs := []struct {
		name    string
		version string
	}{
		{tstName, tstVersion},
		{"yourpkg", "0.0.1"},
	}

	for _, pkg := range pkgs {
		p, err := NewSwPkg("apt", pkg.name, pkg.version)
		assert.NoError(err)
		listPkgs = append(listPkgs, p)
		queryPkgs = append(queryPkgs, p)
	}

	mpkgMgr := &MockPkgManager{
		listPkgs:  listPkgs,
		queryPkgs: queryPkgs,
	}

	mPkg := &MockPkg{
		manager: mpkgMgr,
		name:    tstName,
		version: tstVersion,
	}

	assert.NoError(IsInstalled(mPkg))
}

// TODO: Need to figure out this one
// Might need to allow mock pkg manager
func TestListInstalled(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(nil)
}
