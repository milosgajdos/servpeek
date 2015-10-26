package commander

import "github.com/milosgajdos83/servpeek/utils/command"

const (
	gem = "gem"
)

var (
	// cli arguments to gem
	gemListPkgsArgs = []string{"list", "--local"}
	gemQueryPkgArgs = []string{"list", "--local"}
)

// GemCommander provides gem command manager commands
type GemCommander struct {
	*PkgCommander
}

// NewGemCommander returns gem command manager
func NewGemCommander() *PkgCommander {
	return &PkgCommander{
		ListPkgs: command.NewCommand(gem, gemListPkgsArgs...),
		QueryPkg: command.NewCommand(gem, gemQueryPkgArgs...),
	}
}
