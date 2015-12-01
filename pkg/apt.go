// build linux

package pkg

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const dpkg = "dpkg-query"

var (
	// cli arguments passed to dpkg-query
	dpkgListPkgsCmdArgs = []string{"-l"}
	dpkgQueryPkgCmdArgs = []string{"-W", "-f '${Status} ${Version}'"}
	// apt package manager parser hints
	aptListPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`^ii`),
		matcher: regexp.MustCompile(`^ii\s+(\S+)\s+(\S+)\s+.*`),
	}
	aptQueryPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`^\s+'[A-Za-z]`),
		matcher: regexp.MustCompile(`^\s+'\w+\s+\w+\s+\w+\s+(\S+)'$`),
	}
)

// aptManager implements Apt package manager
type aptManager struct {
	baseManager
}

// NewAptManager returns apt Manager or fails with error
func NewAptManager() Manager {
	return &aptManager{
		baseManager: baseManager{
			Commander: NewAptCommander(),
			pkgType:   "apt",
		},
	}
}

// AptCommander provides aptitude command manager commands
type aptCommander struct {
	baseCommander
}

// NewAptCommander returns Commander that provides apt package manager commands
func NewAptCommander() Commander {
	return &aptCommander{
		baseCommander: baseCommander{
			ListPkgsCmd: command.NewCommand(dpkg, dpkgListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(dpkg, dpkgQueryPkgCmdArgs...),
		},
	}
}

// aptParser provides parser for list and query package manager commands
type aptParser struct {
	baseCmdOutParser
}

// NewAptParser returs CmdOutParser that parses aptitude Manager commands outputs
func NewAptParser() CmdOutParser {
	return &aptParser{
		baseCmdOutParser: baseCmdOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("apt", "list", aptListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("apt", "query", aptQueryPkgsOutHints),
		},
	}
}
