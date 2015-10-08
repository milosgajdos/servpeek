package commander

const (
	rpm = "rpm"
)

var (
	// cli arguments passed to rpm
	rpmListPkgsArgs  = []string{"-qa --qf '%{NAME}%20{VERSION}-%{RELEASE}\n'"}
	rpmQueryPkgsArgs = []string{"-qi"}
)

type YumCommander struct {
	*Commander
}

// yum manager commands
func YumCommander() *Commander {
	return &Commander{
		ListPkgs:  BuildCmd(rpm, rpmListPkgsArgs),
		QueryPkgs: BuildCmd(rpm, rpmQueryPkgsArgs),
	}
}
