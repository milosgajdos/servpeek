package commander

const (
	pip = "pip"
)

var (
	// cli arguments passed to gem
	pipListPkgsArgs  = []string{"list"}
	pipQueryPkgsArgs = []string{"show"}
)

// pip manager commands
func PipCommander() *Commander {
	return &Commander{
		ListPkgs:  BuildCmd(pip, pipListPkgsArgs),
		QueryPkgs: BuildCmd(pip, pipQueryPkgsArgs),
	}
}
