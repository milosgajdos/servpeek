package commander

const (
	rpm = "rpm"
)

var (
	// cli arguments passed to rpm
	rpmListPkgsArgs = []string{"-qa --qf '%{NAME}%20{VERSION}-%{RELEASE}\n'"}
	rpmQueryPkgArgs = []string{"-qi"}
)

// YumCommander provides yum command manager commands
type YumCommander struct {
	*PkgCommander
}

// NewYumCommander returns yum command manager
func NewYumCommander() *PkgCommander {
	return &PkgCommander{
		ListPkgs: BuildCmd(rpm, rpmListPkgsArgs...),
		QueryPkg: BuildCmd(rpm, rpmQueryPkgArgs...),
	}
}
