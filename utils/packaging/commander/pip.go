package commander

import "github.com/milosgajdos83/servpeek/utils/command"

const (
	pip = "pip"
)

var (
	// cli arguments passed to gem
	pipListPkgsArgs = []string{"list"}
	pipQueryPkgArgs = []string{"show"}
)

// PipCommand provides gem command manager commands
type PipCommand struct {
	*PkgCommander
}

// NewPipCommander returns pip command manager
func NewPipCommander() *PkgCommander {
	return &PkgCommander{
		ListPkgs: command.NewCommand(pip, pipListPkgsArgs...),
		QueryPkg: command.NewCommand(pip, pipQueryPkgArgs...),
	}
}
