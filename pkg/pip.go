package pkg

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const pip = "pip"

var (
	// cli arguments passed to gem
	pipListPkgsCmdArgs = []string{"list"}
	pipQueryPkgCmdArgs = []string{"show"}
	// pip package manager parser hints
	pipListPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`^[A-Za-z]`),
		matcher: regexp.MustCompile(`^(\S+)\s+\((\S+)\)$`),
	}
	pipQueryPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`^Version`),
		matcher: regexp.MustCompile(`^Version\s?:\s+(\S+).*`),
	}
)

// pipManager implements pip package manager
type pipManager struct {
	basePkgManager
}

// NewPipManager returns pip PkgManager or fails with error
func NewPipManager() (PkgManager, error) {
	return &pipManager{
		basePkgManager: basePkgManager{
			PkgCommander: NewPipCommander(),
			pkgType:      "pip",
		},
	}, nil
}

// pipCommander provides gem command manager commands
type pipCommander struct {
	basePkgCommander
}

// NewPipCommander returns PkgCommander that provides pip package manager commands
func NewPipCommander() PkgCommander {
	return &pipCommander{
		basePkgCommander: basePkgCommander{
			ListPkgsCmd: command.NewCommand(pip, pipListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(pip, pipQueryPkgCmdArgs...),
		},
	}
}

// pipParser provides parser for list and query package manager commands
type pipParser struct {
	basePkgOutParser
}

// NewPipParser returns PkgOutParser that parses pip PkgManager commands outputs
func NewPipParser() PkgOutParser {
	return &pipParser{
		basePkgOutParser: basePkgOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("pip", "list", pipListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("pip", "query", pipQueryPkgsOutHints),
		},
	}
}
