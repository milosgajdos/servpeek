package manager

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const (
	DpkgQuery = "dpkg-query"
)

var aptCmds = Commands{
	QueryAllInstalled: command.Build(DpkgQuery, "-l"),
	QueryPkgInfo:      command.Build(DpkgQuery, "-W -f '${Status} ${Version}' %s"),
}

// AptManger embeds PkgManager and implements Manager interface
type AptManager struct {
	BasePkgManager
}

// NewAptManager returns Manager or fails with error
func NewAptManager() (PkgManager, error) {
	return &AptManager{
		BasePkgManager: BasePkgManager{
			cmds: aptCmds,
		},
	}, nil
}

func (am *AptManager) QueryAllInstalled() ([]*PkgInfo, error) {
	pkgs := make([]*PkgInfo, 0)
	out, err := command.Run(aptCmds.QueryAllInstalled)
	if err != nil {
		return nil, err
	}

	// TODO: parse dpkg output
	fmt.Printf("%v", out)
	return pkgs, nil
}

func (am *AptManager) QueryInstalled(pkgName ...string) ([]*PkgInfo, error) {
	pkgs := make([]*PkgInfo, 0)
	for _, name := range pkgName {
		aptCmd := fmt.Sprintf(aptCmds.QueryPkgInfo, name)
		out, err := command.Run(aptCmd)
		if err != nil {
			return nil, err
		}
		// TODO: parse dpkg output
		fmt.Printf("%v", out)
	}
	return pkgs, nil
}
