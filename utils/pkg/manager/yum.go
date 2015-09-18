package manager

import (
	"fmt"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const (
	RpmQuery = "rpm"
)

var (
	// cli flags passed to rpm
	rpmQueryAllArgs     = []string{"-qa --qf '%{NAME}%20{VERSION}-%{RELEASE}\n'"}
	rpmQueryPkgInfoArgs = []string{"-qi"}
	// commands provided by yum package manager
	yum = Commands{
		QueryAllInstalled: command.NewCommand(RpmQuery, rpmQueryAllArgs...),
		QueryPkgInfo:      command.NewCommand(RpmQuery, rpmQueryPkgInfoArgs...),
	}
)

// AptManger embeds PkgManager and implements Manager interface
type YumManager struct {
	BasePkgManager
}

// NewYumManager returns PkgManager or fails with error
func NewYumManager() (PkgManager, error) {
	return &YumManager{
		BasePkgManager: BasePkgManager{
			cmds: yum,
		},
	}, nil
}

func (ym YumManager) QueryAllInstalled() ([]*PkgInfo, error) {
	pkgInfos := make([]*PkgInfo, 0)
	cmdOut := ym.cmds.QueryAllInstalled.Run()
	defer cmdOut.Close()

	for cmdOut.Next() {
		line := cmdOut.Text()
		if strings.HasPrefix(line, "") {
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

func (ym *YumManager) QueryInstalled(pkgName ...string) ([]*PkgInfo, error) {
	pkgInfos := make([]*PkgInfo, 0)
	for _, name := range pkgName {
		ym.cmds.QueryPkgInfo.Args = append(ym.cmds.QueryPkgInfo.Args, name)
		cmdOut := ym.cmds.QueryPkgInfo.Run()
		defer cmdOut.Close()

		for cmdOut.Next() {
			line := cmdOut.Text()
			if strings.HasPrefix(line, "Version") {
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

func parseRpmInstalledOut(line string) (*PkgInfo, error) {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return nil, fmt.Errorf("Could not parse rpm info: %s", line)
	}

	return &PkgInfo{
		Name:    fields[0],
		Version: fields[1],
	}, nil
}

func parseRpmInfoOut(line string) (*PkgInfo, error) {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return nil, fmt.Errorf("Could not parse dpkg info: %s", line)
	}

	return &PkgInfo{
		Version: fields[2],
	}, nil
}
