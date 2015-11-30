// build linux

package commander

import "github.com/milosgajdos83/servpeek/utils/command"

const rpm = "rpm"

var (
	// cli arguments passed to rpm
	rpmListPkgsCmdArgs = []string{"-qa --qf '%{NAME}%20{VERSION}-%{RELEASE}\n'"}
	rpmQueryPkgCmdArgs = []string{"-qi"}
)

// YumCommander provides yum command manager commands
type YumCommander struct {
	*BasePkgCommander
}

// NewYumCommander returns PkgCommander that provides yum package manager commands
func NewYumCommander() PkgCommander {
	return &BasePkgCommander{
		ListPkgsCmd: command.NewCommand(rpm, rpmListPkgsCmdArgs...),
		QueryPkgCmd: command.NewCommand(rpm, rpmQueryPkgCmdArgs...),
	}
}
