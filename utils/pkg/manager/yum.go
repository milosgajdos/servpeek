package manager

import "github.com/milosgajdos83/servpeek/utils/command"

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

	// TODO: parse cmdOut rpm output
	return pkgInfos, nil
}

func (ym *YumManager) QueryInstalled(pkgName ...string) ([]*PkgInfo, error) {
	pkgInfos := make([]*PkgInfo, 0)
	for _, name := range pkgName {
		ym.cmds.QueryPkgInfo.Args = append(ym.cmds.QueryPkgInfo.Args, name)
		cmdOut := ym.cmds.QueryPkgInfo.Run()
		defer cmdOut.Close()

		// TODO: parse cmdOut rpm output
	}
	return pkgInfos, nil
}

func parseRpmInstalledOut(line string) (*PkgInfo, error) {
	return nil, nil
}

func parseRpmInfoOut(line string) (*PkgInfo, error) {
	return nil, nil
}
