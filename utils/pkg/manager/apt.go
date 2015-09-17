package manager

import (
	"fmt"
	"strings"

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
	pkgInfos := make([]*PkgInfo, 0)
	out := apt.QueryAllInstalled.Run()
	defer out.Close()

	for out.Next() {
		line := out.Text()
		if strings.HasPrefix(line, "ii") {
			pkgInfo, err := parseDpkgInstalledOut(line)
			if err != nil {
				return nil, err
			}
			pkgInfos = append(pkgInfos, pkgInfo)
		}
	}

	if err := out.Err(); err != nil {
		return nil, err
	}

	return pkgInfos, nil
}

func (am *AptManager) QueryInstalled(pkgName ...string) ([]*PkgInfo, error) {
	pkgInfos := make([]*PkgInfo, 0)
	for _, name := range pkgName {
		pkgInfoArgs := fmt.Sprintf(dpkgQueryPkgInfoArgs, name)
		apt.QueryPkgInfo = command.NewCommand(DpkgQuery, pkgInfoArgs)
		out := apt.QueryPkgInfo.Run()
		defer out.Close()

		for out.Next() {
			line := out.Text()
			pkgInfo, err := parseDpkgInfoOut(line)
			if err != nil {
				return nil, err
			}
			pkgInfo.Name = name
			pkgInfos = append(pkgInfos, pkgInfo)
		}

		if err := out.Err(); err != nil {
			return nil, err
		}
	}
	return pkgInfos, nil
}

func parseDpkgInstalledOut(line string) (*PkgInfo, error) {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return nil, fmt.Errorf("Could not parse dpkg info: %s", line)
	}

	return &PkgInfo{
		Name:    fields[1],
		Version: fields[2],
	}, nil
}

func parseDpkgInfoOut(line string) (*PkgInfo, error) {
	fields := strings.Fields(line)
	if len(fields) < 4 {
		return nil, fmt.Errorf("Could not parse dpkg info: %s", line)
	}

	return &PkgInfo{
		Version: fields[3],
	}, nil
}
