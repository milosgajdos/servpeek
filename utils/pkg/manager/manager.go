// package manager provides an implementation of a software
// package manager. Description sounds kind of silly redundant
// due to unfortunate naming convention classh in Go :-)
package manager

import (
	"fmt"
	"strings"

	"github.com/milosgajdos83/servpeek/utils"
	"github.com/milosgajdos83/servpeek/utils/command"
	"github.com/milosgajdos83/servpeek/utils/pkg/manager/apt"
	"github.com/milosgajdos83/servpeek/utils/pkg/manager/yum"
)

// PkgManager defines package manager interface
type PkgManager interface {
	// QueryInstalledAll returns a slice of all isntalled packages
	ListPkgs() ([]*PkgInfo, error)
	// QueryInstalled returns a list of requested packages
	QueryPkgs(pkgName ...string) ([]*PkgInfo, error)
}

// Commands defines package manager commands
type Commands struct {
	// Query all Installed packages
	ListPkgs *command.Command
	// QueryInstalled queries installed packages
	QueryPkgs *command.Command
}

// PkgInfo contains package name and version
type PkgInfo struct {
	Name    string
	Version string
}

// BasePkgManager is a package manager that implements
// basic package manager comands
type BasePkgManager struct {
	// cmds provides package manager commands
	cmds Commands
	// parse hints are ugly hack to help with parsing cmd output
	parseHints *utils.ParseHints
}

// NewPkgManager returns PkgManager based on the package type
func NewPkgManager(pkgType string) (PkgManager, error) {
	switch pkgType {
	case "apt", "dpkg":
		return NewAptManager()
	case "rpm", "yum":
		return NewYumManager()
	}
	return nil, fmt.Errorf("Unsupported package type")
}

// AptManager implements Apt package manager
type AptManager struct {
	BasePkgManager
}

// NewAptManager returns PkgManager or fails with error
func NewAptManager() (PkgManager, error) {
	return &AptManager{
		BasePkgManager: BasePkgManager{
			cmds: Commands{
				ListPkgs:  utils.BuildCmd(apt.QueryCmd, apt.ListPkgsArgs...),
				QueryPkgs: utils.BuildCmd(apt.QueryCmd, apt.QueryPkgsArgs...),
			},
			parseHints: apt.ParseHints,
		},
	}, nil
}

// YumManager implements Yum package manager
type YumManager struct {
	BasePkgManager
}

// NewYumManager returns PkgManager or fails with error
func NewYumManager() (PkgManager, error) {
	return &YumManager{
		BasePkgManager: BasePkgManager{
			cmds: Commands{
				ListPkgs:  utils.BuildCmd(yum.QueryCmd, yum.ListPkgsArgs...),
				QueryPkgs: utils.BuildCmd(yum.QueryCmd, yum.QueryPkgsArgs...),
			},
			parseHints: yum.ParseHints,
		},
	}, nil
}

// ListPkgs queries packages manager for installed packages and returns then in slice
// It fails if either the command fails or command output parser fails
func (pm *BasePkgManager) ListPkgs() ([]*PkgInfo, error) {
	pkgInfos := make([]*PkgInfo, 0)
	cmdOut := pm.cmds.ListPkgs.Run()
	defer cmdOut.Close()

	for cmdOut.Next() {
		line := cmdOut.Text()
		if strings.HasPrefix(line, pm.parseHints.ListPrefix) {
			pkgInfo, err := parseInstalledOut(line, pm.parseHints)
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

// QueryPkgs queries package manager for information about an arbitrary list of packages
// It fails if either the command query fails or command output parser fails
func (pm *BasePkgManager) QueryPkgs(pkgName ...string) ([]*PkgInfo, error) {
	pkgInfos := make([]*PkgInfo, 0)
	for _, name := range pkgName {
		pm.cmds.QueryPkgs.Args = append(pm.cmds.QueryPkgs.Args, name)
		cmdOut := pm.cmds.QueryPkgs.Run()
		defer cmdOut.Close()

		for cmdOut.Next() {
			line := cmdOut.Text()
			if strings.HasPrefix(line, pm.parseHints.QueryPrefix) {
				pkgInfo, err := parseQueryOut(line, pm.parseHints)
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

func parseQueryOut(line string, pHints *utils.ParseHints) (*PkgInfo, error) {
	fields := strings.Fields(line)
	if len(fields) < pHints.ListMinFields {
		return nil, fmt.Errorf("Could not parse package info: %s", line)
	}

	return &PkgInfo{
		Name:    fields[pHints.ListVersionIdx-1],
		Version: fields[pHints.ListVersionIdx],
	}, nil
}

func parseInstalledOut(line string, pHints *utils.ParseHints) (*PkgInfo, error) {
	fields := strings.Fields(line)
	if len(fields) < pHints.QueryMinFields {
		return nil, fmt.Errorf("Could not parse package info: %s", line)
	}

	return &PkgInfo{
		Version: fields[pHints.QueryVersionIdx],
	}, nil
}
