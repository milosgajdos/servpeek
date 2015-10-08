package commander

const (
	gem = "gem"
)

var (
	// cli arguments to gem
	gemListPkgsArgs = []string{"list", "--local"}
	gemQueryPkgArgs = []string{"list", "--local"}
)

type GemCommander struct {
	*Commander
}

// gem manager commands
func NewGemCommander() *Commander {
	return &Commander{
		ListPkgs: BuildCmd(gem, gemListPkgsArgs...),
		QueryPkg: BuildCmd(gem, gemQueryPkgArgs...),
	}
}
