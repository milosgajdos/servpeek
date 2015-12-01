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
	basePkgManager
}

// NewAptManager returns apt PkgManager or fails with error
func NewAptManager() (PkgManager, error) {
	return &aptManager{
		basePkgManager: basePkgManager{
			PkgCommander: NewAptCommander(),
			pkgType:      "apt",
		},
	}, nil
}

// AptCommander provides aptitude command manager commands
type aptCommander struct {
	basePkgCommander
}

// NewAptCommander returns PkgCommander that provides apt package manager commands
func NewAptCommander() PkgCommander {
	return &aptCommander{
		basePkgCommander: basePkgCommander{
			ListPkgsCmd: command.NewCommand(dpkg, dpkgListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(dpkg, dpkgQueryPkgCmdArgs...),
		},
	}
}

// aptParser provides parser for list and query package manager commands
type aptParser struct {
	basePkgOutParser
}

// NewAptParser returs PkgOutParser that parses aptitude PkgManager commands outputs
func NewAptParser() PkgOutParser {
	return &aptParser{
		basePkgOutParser: basePkgOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("apt", "list", aptListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("apt", "query", aptQueryPkgsOutHints),
		},
	}
}
