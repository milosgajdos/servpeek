package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkg(t *testing.T) {
	assert := assert.New(t)
	p, err := NewSwPkg("gem", "PkgName", "0.1.1")
	assert.NoError(err)

	expected := "[Package] Type: " + p.Type() + " Name: " + p.Name() + " Version: " + p.Version()
	assert.Equal(expected, p.String())
}

func TestType(t *testing.T) {
	assert := assert.New(t)
	p, err := NewSwPkg("gem", "PkgName", "0.1.1")
	assert.NoError(err)
	assert.Equal("gem", p.Type())
}

func TestName(t *testing.T) {
	assert := assert.New(t)
	p, err := NewSwPkg("pip", "PkgName", "0.1.1")
	assert.NoError(err)
	assert.Equal("PkgName", p.Name())
}

func TestVersion(t *testing.T) {
	assert := assert.New(t)
	p, err := NewSwPkg("pip", "PkgName", "0.1.1")
	assert.NoError(err)
	assert.Equal("0.1.1", p.Version())
}
