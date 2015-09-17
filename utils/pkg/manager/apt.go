package manager

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const (
	DpkgQuery            = "dpkg-query"
	dpkgQueryAllArgs     = "-l"
	dpkgQueryPkgInfoArgs = "-W -f '${Status} ${Version}' %s"
)

var apt = Commands{
	QueryAllInstalled: command.NewCommand(DpkgQuery, dpkgQueryAllArgs),
}

// AptManger embeds PkgManager and implements Manager interface
type AptManager struct {
	BasePkgManager
}

// NewAptManager returns Manager or fails with error
func NewAptManager() (PkgManager, error) {
	return &AptManager{
		BasePkgManager: BasePkgManager{
			cmds: apt,
		},
	}, nil
}

func (am *AptManager) QueryAllInstalled() ([]*PkgInfo, error) {
	pkgs := make([]*PkgInfo, 0)
	out := apt.QueryAllInstalled.Run()

	// TODO: parse dpkg output
	fmt.Printf("%v", out)
	return pkgs, nil
}

func (am *AptManager) QueryInstalled(pkgName ...string) ([]*PkgInfo, error) {
	pkgs := make([]*PkgInfo, 0)
	for _, name := range pkgName {
		pkgInfoArgs := fmt.Sprintf(dpkgQueryPkgInfoArgs, name)
		apt.QueryPkgInfo = command.NewCommand(DpkgQuery, pkgInfoArgs)
		out := apt.QueryPkgInfo.Run()

		// TODO: parse dpkg output
		fmt.Printf("%v", out)
	}
	return pkgs, nil
}

func parseDpkgInstalledOut(line string) (*PkgInfo, error) {
	return nil, nil
}

func parseDpkgInfoOut(line string) (*PkgInfo, error) {
	return nil, nil
}
