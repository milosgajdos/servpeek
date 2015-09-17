package manager

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const (
	RpmQuery            = "rpm"
	rpmQueryAllArgs     = "-qa --qf '%{NAME}%20{VERSION}-%{RELEASE}\n'"
	rpmQueryPkgInfoArgs = "-qi %s"
)

var yum = Commands{
	QueryAllInstalled: command.NewCommand(RpmQuery, rpmQueryAllArgs),
}

// AptManger embeds PkgManager and implements Manager interface
type YumManager struct {
	BasePkgManager
}

func NewYumManager() (PkgManager, error) {
	return &YumManager{
		BasePkgManager: BasePkgManager{
			cmds: yum,
		},
	}, nil
}

func (ym YumManager) QueryAllInstalled() ([]*PkgInfo, error) {
	pkgs := make([]*PkgInfo, 0)
	out := yum.QueryAllInstalled.Run()

	// TODO: parse dpkg output
	fmt.Printf("%v", out)
	return pkgs, nil
}

func (ym *YumManager) QueryInstalled(pkgName ...string) ([]*PkgInfo, error) {
	pkgs := make([]*PkgInfo, 0)
	for _, name := range pkgName {
		pkgInfoArgs := fmt.Sprintf(rpmQueryPkgInfoArgs, name)
		yum.QueryPkgInfo = command.NewCommand(DpkgQuery, pkgInfoArgs)
		out := yum.QueryPkgInfo.Run()

		// TODO: parse dpkg output
		fmt.Printf("%v", out)
	}
	return pkgs, nil
}

func parseRpmInstalledOut(line string) (*PkgInfo, error) {
	return nil, nil
}

func parseRpmInfoOut(line string) (*PkgInfo, error) {
	return nil, nil
}
