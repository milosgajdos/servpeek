package commander

const (
	gem = "gem"
)

var (
	// cli arguments to gem
	gemListPkgsArgs  = []string{"list", "--local"}
	gemQueryPkgsArgs = []string{"list", "--local"}
)

type GemCommander struct {
	*Commander
}

// gem manager commands
func GemCommander() *Commander {
	return &Commander{
		ListPkgs:  BuildCmd(gem, gemListPkgsArgs),
		QueryPkgs: BuildCmd(gem, gemQueryPkgsArgs),
	}
}
