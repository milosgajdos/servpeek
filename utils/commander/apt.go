package commander

const (
	dpkg = "dpkg-query"
)

var (
	// cli arguments passed to dpkg-query
	dpkgListPkgsArgs = []string{"-l"}
	dpkgQueryPkgArgs = []string{"-W", "-f '${Status} ${Version}'"}
)

type AptCommander struct {
	*Commander
}

// aptitude manager commands
func NewAptCommander() *Commander {
	return &Commander{
		ListPkgs: BuildCmd(dpkg, dpkgListPkgsArgs...),
		QueryPkg: BuildCmd(dpkg, dpkgQueryPkgArgs...),
	}
}
