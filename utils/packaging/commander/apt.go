// build linux

package commander

import "github.com/milosgajdos83/servpeek/utils/command"

const dpkg = "dpkg-query"

var (
	// cli arguments passed to dpkg-query
	dpkgListPkgsCmdArgs = []string{"-l"}
	dpkgQueryPkgCmdArgs = []string{"-W", "-f '${Status} ${Version}'"}
)

// AptCommander provides aptitude command manager commands
type AptCommander struct {
	*BasePkgCommander
}

// NewAptCommander returns PkgCommander that provides apt package manager commands
func NewAptCommander() PkgCommander {
	return &BasePkgCommander{
		ListPkgsCmd: command.NewCommand(dpkg, dpkgListPkgsCmdArgs...),
		QueryPkgCmd: command.NewCommand(dpkg, dpkgQueryPkgCmdArgs...),
	}
}
