package pkg

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const gem = "gem"

var (
	// cli arguments to gem
	gemListPkgsCmdArgs = []string{"list", "--local"}
	gemQueryPkgCmdArgs = []string{"list", "--local"}
	// gem package manager parser hints
	gemListPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`^[A-Za-z]`),
		matcher: regexp.MustCompile(`^(\S+)\s+\((\S+)\)$`),
	}
	gemQueryPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`^[A-Za-z]`),
		matcher: regexp.MustCompile(`^\S+\s+\((\S+)\)$`),
	}
)

// GemManager implements gem package manager
type gemManager struct {
	baseManager
}

// NewGemManager returns gem Manager or fails with error
func NewGemManager() Manager {
	return &gemManager{
		baseManager: baseManager{
			Commander: NewGemCommander(),
			pkgType:   "gem",
		},
	}
}

// gemCommander provides gem command manager commands
type gemCommander struct {
	baseCommander
}

// NewGemCommander returns Commander that provides gem package manager commands
func NewGemCommander() Commander {
	return &gemCommander{
		baseCommander: baseCommander{
			ListPkgsCmd: command.NewCommand(gem, gemListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(gem, gemQueryPkgCmdArgs...),
		},
	}
}

// gemParser provides parser for list and query package manager commands
type gemParser struct {
	baseCmdOutParser
}

// NewGemParser returns CmdOutParser that parses gem Manager commands outputs
func NewGemParser() CmdOutParser {
	return &gemParser{
		baseCmdOutParser: baseCmdOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("gem", "list", gemListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("gem", "query", gemQueryPkgsOutHints),
		},
	}
}
