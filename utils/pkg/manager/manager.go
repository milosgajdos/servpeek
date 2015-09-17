// package manager provides an implementation of package manager
// this sounds kind of silly redundant, so ignore the naming
package manager

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// PkgManager defines package manager interface
type PkgManager interface {
	// QueryInstalledAll returns a slice of all isntalled packages
	QueryAllInstalled() ([]*PkgInfo, error)
	// QueryInstalled returns a list of requested packages
	QueryInstalled(pkgName ...string) ([]*PkgInfo, error)
}

// Commands defines package manager commands
type Commands struct {
	// Query all Installed packages
	QueryAllInstalled *command.Command
	// QueryInstalled queries installed packages
	QueryPkgInfo *command.Command
}

// PkgInfo
type PkgInfo struct {
	Name    string
	Version string
}

// BasePkgManager is a baic package Manager that has a package manager
// with particular package manager comands
type BasePkgManager struct {
	// cmds provides package manager commands
	cmds Commands
}

// NewPkgManager returns Manager based on the package type it manages
func NewPkgManager(pkgType string) (PkgManager, error) {
	switch pkgType {
	case "apt", "dpkg":
		return NewAptManager()
	case "rpm", "yum":
		return NewYumManager()
	}

	return nil, fmt.Errorf("Unsupported package type")
}
