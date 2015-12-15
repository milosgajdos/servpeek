package pkg

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInstalled(t *testing.T) {
	assert := assert.New(t)
	parsers := map[string]CmdOutParser{
		"apt": NewAptParser(),
		"yum": NewYumParser(),
		"apk": NewApkParser(),
		"pip": NewPipParser(),
		"gem": NewGemParser(),
	}

	pkgs := []struct {
		pkgType string
		name    string
		version string
	}{
		{"apt", "cron", "3.0pl1-124ubuntu2"},
		{"yum", "grep", "2.20"},
		{"apk", "musl-utils", "1.1.11-r2"},
		{"pip", "setuptools", ""},
		{"gem", "bundler", "1.10.6"},
	}

	for _, pkg := range pkgs {
		currentDir, err := os.Getwd()
		assert.NoError(err)
		fixturesPath := path.Join(currentDir, "test-fixtures", pkg.pkgType+"query.out")
		cmdOut, err := ioutil.ReadFile(fixturesPath)
		assert.NoError(err)
		p := &mockPkg{
			manager: &mockManager{
				queryCmd: &mockCommander{
					cmdOut: string(cmdOut),
				},
				parser:  parsers[pkg.pkgType],
				pkgType: pkg.pkgType,
			},
			name:    pkg.name,
			version: pkg.version,
		}
		assert.NoError(IsInstalled(p))
	}

	// check package that is not installed
	p := &mockPkg{
		manager: &mockManager{
			queryCmd: &mockCommander{
				cmdOut: "\n*** LOCAL GEMS ***\n\n\n",
			},
			parser:  NewGemParser(),
			pkgType: "gem",
		},
		name:    "randpkg",
		version: "",
	}
	assert.Error(IsInstalled(p))
}

func TestListPkgs(t *testing.T) {
	assert := assert.New(t)
	// parse list pkgs output
	currentDir, err := os.Getwd()
	assert.NoError(err)
	fixturesPath := path.Join(currentDir, "test-fixtures", "gemlist.out")
	gemListOut, err := ioutil.ReadFile(fixturesPath)
	pkgMgr := &mockManager{
		listCmd: &mockCommander{
			cmdOut: string(gemListOut),
		},
		parser:  NewGemParser(),
		pkgType: "gem",
	}

	pkgs, err := pkgMgr.ListPkgs()
	assert.NoError(err)
	assert.NotEmpty(pkgs)

	// can't parse output from listpkgs
	pkgMgr = &mockManager{
		listCmd: &mockCommander{
			cmdOut: "garbage",
		},
		parser:  NewGemParser(),
		pkgType: "gem",
	}

	pkgs, err = pkgMgr.ListPkgs()
	assert.Error(err)
	assert.Empty(pkgs)
}
