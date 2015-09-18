package manager

import (
	"fmt"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const (
	DpkgQuery = "dpkg-query"
)

var (
	// cli flags passed to dpkg-query
	dpkgQueryAllArgs     = []string{"-l"}
	dpkgQueryPkgInfoArgs = []string{"-W", "-f '${Status} ${Version}'"}
)

// AptManger embeds PkgManager and implements Manager interface
type AptManager struct {
	BasePkgManager
}

// NewAptManager returns PkgManager or fails with error
func NewAptManager() (PkgManager, error) {
	// commands provided by apt package manager
	apt := Commands{
		QueryAllInstalled: command.NewCommand(DpkgQuery, dpkgQueryAllArgs...),
		QueryPkgInfo:      command.NewCommand(DpkgQuery, dpkgQueryPkgInfoArgs...),
	}

	return &AptManager{
		BasePkgManager: BasePkgManager{
			cmds: apt,
		},
	}, nil
}

func (am *AptManager) QueryAllInstalled() ([]*PkgInfo, error) {
	pkgInfos := make([]*PkgInfo, 0)
	cmdOut := am.cmds.QueryAllInstalled.Run()
	defer cmdOut.Close()

	for cmdOut.Next() {
		line := cmdOut.Text()
		if strings.HasPrefix(line, "ii") {
			pkgInfo, err := parseDpkgInstalledOut(line)
			if err != nil {
				return nil, err
			}
			pkgInfos = append(pkgInfos, pkgInfo)
		}
	}

	if err := cmdOut.Err(); err != nil {
		return nil, err
	}

	return pkgInfos, nil
}

func (am *AptManager) QueryInstalled(pkgName ...string) ([]*PkgInfo, error) {
	pkgInfos := make([]*PkgInfo, 0)
	for _, name := range pkgName {
		am.cmds.QueryPkgInfo.Args = append(am.cmds.QueryPkgInfo.Args, name)
		cmdOut := am.cmds.QueryPkgInfo.Run()
		defer cmdOut.Close()

		for cmdOut.Next() {
			line := cmdOut.Text()
			if strings.HasPrefix(line, "") {
				pkgInfo, err := parseDpkgInfoOut(line)
				if err != nil {
					return nil, err
				}
				pkgInfo.Name = name
				pkgInfos = append(pkgInfos, pkgInfo)
			}
		}

		if err := cmdOut.Err(); err != nil {
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
