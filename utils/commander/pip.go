package commander

const (
	pip = "pip"
)

var (
	// cli arguments passed to gem
	pipListPkgsArgs = []string{"list"}
	pipQueryPkgArgs = []string{"show"}
)

type PipCommand struct {
	*Commander
}

// pip manager commands
func NewPipCommander() *Commander {
	return &Commander{
		ListPkgs: BuildCmd(pip, pipListPkgsArgs...),
		QueryPkg: BuildCmd(pip, pipQueryPkgArgs...),
	}
}
