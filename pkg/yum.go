// build linux

package pkg

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const rpm = "rpm"

var (
	// cli arguments passed to rpm
	rpmListPkgsCmdArgs = []string{"-qa --qf '%{NAME}%20{VERSION}-%{RELEASE}\n'"}
	rpmQueryPkgCmdArgs = []string{"-qi"}
	// yum package manager parser hints
	rpmListPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`^[A-Za-z]`),
		matcher: regexp.MustCompile(`^(\S+)\s+(\S+).*`),
	}
	rpmQueryPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`^Version`),
		matcher: regexp.MustCompile(`^Version\s+:\s+(\S+).*`),
	}
)

// yumManager implements Yum package manager
type yumManager struct {
	basePkgManager
}

// NewYumManager returns yum PkgManager or fails with error
func NewYumManager() (PkgManager, error) {
	return &yumManager{
		basePkgManager: basePkgManager{
			PkgCommander: NewYumCommander(),
			pkgType:      "yum",
		},
	}, nil
}

// yumCommander provides yum command manager commands
type yumCommander struct {
	basePkgCommander
}

// NewYumCommander returns PkgCommander that provides yum package manager commands
func NewYumCommander() PkgCommander {
	return &yumCommander{
		basePkgCommander: basePkgCommander{
			ListPkgsCmd: command.NewCommand(rpm, rpmListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(rpm, rpmQueryPkgCmdArgs...),
		},
	}
}

type yumParser struct {
	basePkgOutParser
}

// NewYumParser returs PkgOutParser that parses yum PkgManager commands outputs
func NewYumParser() PkgOutParser {
	return &yumParser{
		basePkgOutParser: basePkgOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("yum", "list", rpmListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("yum", "query", rpmQueryPkgsOutHints),
		},
	}
}
