package manager

import (
	"strings"
)

const (
	DpkgQuery = "dpkg-query"
	AptCache  = "apt-cache"
)

var cmds = Commands{
	QueryInstalled: strings.Join([]string{DpkgQuery, "--get-selections"}, " "),
}

// AptManger embeds PkgManager and implements Manager interface
type AptManager struct {
	PkgManager
}

// NewAptManager returns Manager or fails with error
func NewAptManager() (Manager, error) {
	return &AptManager{
		PkgManager: PkgManager{
			Cmds: cmds,
		},
	}, nil
}

func (am *AptManager) CheckInstalled(name string) bool {
	return false
}

func (am *AptManager) CheckInstalledVersion(name string, version string) bool {
	return false
}
