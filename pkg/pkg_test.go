package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkg(t *testing.T) {
	assert := assert.New(t)
	p, err := NewSwPkg("gem", "PkgName", "0.1.1")
	assert.NoError(err)

	expected := "[Pkg] Type: " + p.Manager().Type() + " Name: " + p.Name() + " Version: " + p.Version()
	assert.Equal(expected, p.String())
}

func TestSwPkgManagerType(t *testing.T) {
	assert := assert.New(t)
	ptypes := []string{"apt", "apk", "yum", "pip", "gem"}
	for _, ptype := range ptypes {
		p, err := NewSwPkg(ptype, "PkgName", "0.1.1")
		assert.NoError(err)
		assert.Equal(ptype, p.Manager().Type())
	}
	// Unsupported package manager
	_, err := NewSwPkg("random", "PkgName", "0.1.1")
	assert.Error(err)
}

func TestName(t *testing.T) {
	tstName := "PkgName"
	assert := assert.New(t)
	p, err := NewSwPkg("pip", tstName, "")
	assert.NoError(err)
	assert.Equal(tstName, p.Name())

	// Name can not be empty
	_, err = NewSwPkg("pip", "", "")
	assert.Error(err)
}

func TestVersion(t *testing.T) {
	tstVersion := "0.1.1"
	assert := assert.New(t)
	p, err := NewSwPkg("pip", "PkgName", tstVersion)
	assert.NoError(err)
	assert.Equal(tstVersion, p.Version())
}
