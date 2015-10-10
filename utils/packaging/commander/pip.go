package commander

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
		ListPkgs: BuildCmd(pip, pipListPkgsArgs...),
		QueryPkg: BuildCmd(pip, pipQueryPkgArgs...),
	}
}
