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
	basePkgManager
}

// NewGemManager returns gem PkgManager or fails with error
func NewGemManager() (PkgManager, error) {
	return &gemManager{
		basePkgManager: basePkgManager{
			PkgCommander: NewGemCommander(),
			pkgType:      "gem",
		},
	}, nil
}

// gemCommander provides gem command manager commands
type gemCommander struct {
	basePkgCommander
}

// NewGemCommander returns PkgCommander that provides gem package manager commands
func NewGemCommander() PkgCommander {
	return &gemCommander{
		basePkgCommander: basePkgCommander{
			ListPkgsCmd: command.NewCommand(gem, gemListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(gem, gemQueryPkgCmdArgs...),
		},
	}
}

// gemParser provides parser for list and query package manager commands
type gemParser struct {
	basePkgOutParser
}

// NewGemParser returns PkgOutParser that parses gem PkgManager commands outputs
func NewGemParser() PkgOutParser {
	return &gemParser{
		basePkgOutParser: basePkgOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("gem", "list", gemListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("gem", "query", gemQueryPkgsOutHints),
		},
	}
}
