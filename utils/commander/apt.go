package commander

const (
	dpkg = "dpkg-query"
)

var (
	// cli arguments passed to dpkg-query
	dpkgListPkgsArgs  = []string{"-l"}
	dpkgQueryPkgsArgs = []string{"-W", "-f '${Status} ${Version}'"}
)

type AptCommander struct {
	*Commander
}

// aptitude manager commands
func NewAptCommander() *Commander {
	return &Commander{
		ListPkgs:  BuildCmd(dpkg, dpkgListPkgsArgs),
		QueryPkgs: BuildCmd(dpkg, dpkgQueryPkgsArgs),
	}
}
