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
		filter:  regexp.MustCompile(`^install.*`),
		matcher: regexp.MustCompile(`^install\s+\w+\s+\w+\s+(\S+)$`),
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
			ListPkgsCmd: command.NewCommand(dpkg, dpkgListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(dpkg, dpkgQueryPkgCmdArgs...),
			parser:      NewAptParser(),
			pkgType:     "apt",
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
