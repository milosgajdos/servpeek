package commander

import "github.com/milosgajdos83/servpeek/utils/command"

const (
	dpkg = "dpkg-query"
)

var (
	// cli arguments passed to dpkg-query
	dpkgListPkgsArgs = []string{"-l"}
	dpkgQueryPkgArgs = []string{"-W", "-f '${Status} ${Version}'"}
)

// AptCommander provides aptitude command manager commands
type AptCommander struct {
	*PkgCommander
}

// NewAptCommander returns aptitude command manager
func NewAptCommander() *PkgCommander {
	return &PkgCommander{
		ListPkgs: command.NewCommand(dpkg, dpkgListPkgsArgs...),
		QueryPkg: command.NewCommand(dpkg, dpkgQueryPkgArgs...),
	}
}
