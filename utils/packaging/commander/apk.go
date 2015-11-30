// build linux

package commander

import "github.com/milosgajdos83/servpeek/utils/command"

const apk = "apk"

var (
	// cli arguments passed to dpkg-query
	apkListPkgsCmdArgs = []string{"info", "-v"}
	apkQueryPkgCmdArgs = []string{"info"}
)

// ApkCommander provides apk command manager commands
type ApkCommander struct {
	*BasePkgCommander
}

// NewApkCommander returns PkgCommander that provides apk package manager commands
func NewApkCommander() PkgCommander {
	return &BasePkgCommander{
		ListPkgsCmd: command.NewCommand(apk, apkListPkgsCmdArgs...),
		QueryPkgCmd: command.NewCommand(apk, apkQueryPkgCmdArgs...),
	}
}
