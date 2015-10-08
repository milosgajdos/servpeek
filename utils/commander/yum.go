package commander

const (
	rpm = "rpm"
)

var (
	// cli arguments passed to rpm
	rpmListPkgsArgs = []string{"-qa --qf '%{NAME}%20{VERSION}-%{RELEASE}\n'"}
	rpmQueryPkgArgs = []string{"-qi"}
)

type YumCommander struct {
	*Commander
}

// yum manager commands
func NewYumCommander() *Commander {
	return &Commander{
		ListPkgs: BuildCmd(rpm, rpmListPkgsArgs...),
		QueryPkg: BuildCmd(rpm, rpmQueryPkgArgs...),
	}
}
