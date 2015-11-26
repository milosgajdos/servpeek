package commander

import "github.com/milosgajdos83/servpeek/utils/command"

const gem = "gem"

var (
	// cli arguments to gem
	gemListPkgsCmdArgs = []string{"list", "--local"}
	gemQueryPkgCmdArgs = []string{"list", "--local"}
)

// GemCommander provides gem command manager commands
type GemCommander struct {
	*BaseCommander
}

// NewGemCommander returns PkgCommander that provides gem package manager commands
func NewGemCommander() PkgCommander {
	return &BaseCommander{
		ListPkgsCmd: command.NewCommand(gem, gemListPkgsCmdArgs...),
		QueryPkgCmd: command.NewCommand(gem, gemQueryPkgCmdArgs...),
	}
}
