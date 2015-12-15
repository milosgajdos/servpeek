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
	baseManager
}

// NewPipManager returns pip Manager or fails with error
func NewPipManager() Manager {
	return &pipManager{
		baseManager: baseManager{
			ListPkgsCmd: command.NewCommand(pip, pipListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(pip, pipQueryPkgCmdArgs...),
			parser:      NewPipParser(),
			pkgType:     "pip",
		},
	}
}

// pipParser provides parser for list and query package manager commands
type pipParser struct {
	baseCmdOutParser
}

// NewPipParser returns CmdOutParser that parses pip Manager commands outputs
func NewPipParser() CmdOutParser {
	return &pipParser{
		baseCmdOutParser: baseCmdOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("pip", "list", pipListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("pip", "query", pipQueryPkgsOutHints),
		},
	}
}
