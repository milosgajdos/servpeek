package manager

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const (
	RpmQuery = "rpm"
)

var yumCmds = Commands{
	QueryAllInstalled: command.Build(RpmQuery, "-qa"),
	QueryPkgInfo:      command.Build(RpmQuery, "-qi %s"),
}

// AptManger embeds PkgManager and implements Manager interface
type YumManager struct {
	BasePkgManager
}

func NewYumManager() (PkgManager, error) {
	return &YumManager{
		BasePkgManager: BasePkgManager{
			cmds: yumCmds,
		},
	}, nil
}

func (ym YumManager) QueryAllInstalled() ([]*PkgInfo, error) {
	pkgs := make([]*PkgInfo, 0)
	out, err := command.Run(yumCmds.QueryAllInstalled)
	if err != nil {
		return nil, err
	}

	// TODO: parse rpm output
	fmt.Printf("%v", out)
	return pkgs, nil
}

func (ym *YumManager) QueryInstalled(pkgName ...string) ([]*PkgInfo, error) {
	pkgs := make([]*PkgInfo, 0)
	for _, name := range pkgName {
		yumCmd := fmt.Sprintf(yumCmds.QueryPkgInfo, name)
		out, err := command.Run(yumCmd)
		if err != nil {
			return nil, err
		}
		// TODO: parse rpm output
		fmt.Printf("%v", out)
	}

	return pkgs, nil
}
