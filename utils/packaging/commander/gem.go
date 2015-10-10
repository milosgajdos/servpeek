package commander

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
		ListPkgs: BuildCmd(gem, gemListPkgsArgs...),
		QueryPkg: BuildCmd(gem, gemQueryPkgArgs...),
	}
}
