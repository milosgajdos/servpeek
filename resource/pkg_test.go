package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkg(t *testing.T) {
	assert := assert.New(t)

	p := &Pkg{
		Type:    "apt",
		Name:    "PkgName",
		Version: "0.1.1",
	}

	assert.Equal("[Package] Type: "+p.Type+" Name: "+p.Name+" Version: "+p.Version, p.String())
}
