package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkg(t *testing.T) {
	assert := assert.New(t)
	p, err := NewPackage("gem", "PkgName", "0.1.1")
	assert.NoError(err)

	expected := fmt.Sprintf("[Package] Type: %s Name: %s Version: %v",
		p.Manager().Type(), p.Name(), p.Versions())
	assert.Equal(expected, p.String())
}

func TestPackageManagerType(t *testing.T) {
	assert := assert.New(t)
	for ptype := range supportedPkgTypes {
		p, err := NewPackage(ptype, "PkgName", "0.1.1")
		assert.NoError(err)
		assert.Equal(ptype, p.Manager().Type())
	}
	// Unsupported package manager
	_, err := NewPackage("random", "PkgName", "0.1.1")
	assert.Error(err)
}

func TestName(t *testing.T) {
	tstName := "PkgName"
	assert := assert.New(t)
	p, err := NewPackage("pip", tstName, "")
	assert.NoError(err)
	assert.Equal(tstName, p.Name())

	// Name can not be empty
	_, err = NewPackage("pip", "", "")
	assert.Error(err)
}

func TestVersions(t *testing.T) {
	tstVersion := "0.1.1"
	assert := assert.New(t)
	p, err := NewPackage("pip", "PkgName", tstVersion)
	assert.NoError(err)
	versions := []string{tstVersion}
	assert.Equal(versions, p.Versions())
}
