package commander

import "github.com/milosgajdos83/servpeek/utils/command"

const pip = "pip"

var (
	// cli arguments passed to gem
	pipListPkgsCmdArgs = []string{"list"}
	pipQueryPkgCmdArgs = []string{"show"}
)

// PipCommand provides gem command manager commands
type PipCommand struct {
	*BasePkgCommander
}

// NewPipCommander returns PkgCommander that provides pip package manager commands
func NewPipCommander() PkgCommander {
	return &BasePkgCommander{
		ListPkgsCmd: command.NewCommand(pip, pipListPkgsCmdArgs...),
		QueryPkgCmd: command.NewCommand(pip, pipQueryPkgCmdArgs...),
	}
}
