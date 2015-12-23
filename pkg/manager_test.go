package pkg

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPkgManager(t *testing.T) {
	assert := assert.New(t)
	mgrTypes := []string{"apt", "yum", "apk", "pip", "gem"}
	for _, mgrType := range mgrTypes {
		m, err := NewManager(mgrType)
		assert.NoError(err)
		assert.Equal(mgrType, m.Type())
	}
	// Unsupported package manager
	_, err := NewManager("random")
	assert.Error(err)
}

func TestPkgManagerListPkgs(t *testing.T) {
	assert := assert.New(t)
	parsers := map[string]CmdOutParser{
		"apt": NewAptParser(),
		"yum": NewYumParser(),
		"apk": NewApkParser(),
		"pip": NewPipParser(),
		"gem": NewGemParser(),
	}

	for pkgType, parser := range parsers {
		currentDir, err := os.Getwd()
		assert.NoError(err)
		fixturesPath := path.Join(currentDir, "test-fixtures", pkgType+"list.out")
		cmdOut, err := ioutil.ReadFile(fixturesPath)
		assert.NoError(err)
		mockMgr := &mockPkgManager{
			listCmd: &mockPkgCommand{
				cmdOut: string(cmdOut),
			},
			parser:  parser,
			pkgType: pkgType,
		}

		pkgs, err := mockMgr.ListPkgs()
		assert.NoError(err)
		assert.NotEmpty(pkgs)
	}
}
