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
	baseManager
}

// NewYumManager returns yum Manager or fails with error
func NewYumManager() Manager {
	return &yumManager{
		baseManager: baseManager{
			Commander: NewYumCommander(),
			pkgType:   "yum",
		},
	}
}

// yumCommander provides yum command manager commands
type yumCommander struct {
	baseCommander
}

// NewYumCommander returns Commander that provides yum package manager commands
func NewYumCommander() Commander {
	return &yumCommander{
		baseCommander: baseCommander{
			ListPkgsCmd: command.NewCommand(rpm, rpmListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(rpm, rpmQueryPkgCmdArgs...),
		},
	}
}

type yumParser struct {
	baseCmdOutParser
}

// NewYumParser returs CmdOutParser that parses yum Manager commands outputs
func NewYumParser() CmdOutParser {
	return &yumParser{
		baseCmdOutParser: baseCmdOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("yum", "list", rpmListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("yum", "query", rpmQueryPkgsOutHints),
		},
	}
}
