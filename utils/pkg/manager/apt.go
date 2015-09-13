package manager

import (
	"fmt"
	"regexp"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const (
	DpkgQuery = "dpkg-query"
)

var cmds = Commands{
	QueryInstalled:        command.Build(DpkgQuery, "-s %s"),
	QueryInstalledVersion: command.Build(DpkgQuery, "-W -f '${Status} ${Version}' %s"),
}

// AptManger embeds PkgManager and implements Manager interface
type AptManager struct {
	PkgManager
}

// NewAptManager returns Manager or fails with error
func NewAptManager() (Manager, error) {
	return &AptManager{
		PkgManager: PkgManager{
			cmds: cmds,
		},
	}, nil
}

func (am *AptManager) CheckInstalled(name string) (bool, error) {
	aptCmd := fmt.Sprintf(cmds.QueryInstalled, name)
	_, err := command.Run(aptCmd)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (am *AptManager) CheckInstalledVersion(name string, version string) (bool, error) {
	aptCmd := fmt.Sprintf(cmds.QueryInstalledVersion, name)
	versionRE := fmt.Sprintf("^(install|hold) ok installed %s$", version)
	expectedOut := regexp.MustCompile(versionRE)
	out, err := command.Run(aptCmd)
	if err != nil {
		return false, err
	}

	if match := expectedOut.MatchString(out); match {
		return true, nil
	}

	return false, nil
}
