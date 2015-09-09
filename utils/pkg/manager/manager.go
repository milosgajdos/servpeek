// package manager provides an implementation of package manager
package manager

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

// Manager defines generic software package manager interface
type Manager interface {
	// CheckInstalled returs true if pkgName package is installed
	CheckInstalled(pkgName string) bool
	// CheckInstalledVersion returns true if pkgName package is installed
	// and its version is pkgVersion
	CheckInstalledVersion(pkgName string, pkgVersion string) bool
}

// Commands provides package manager commands
type Commands struct {
	// QueryInstalled queries installed images
	QueryInstalled string
}

// PkgManager is a package Manager that has a list of Commands
type PkgManager struct {
	// Cmds provides package manager commands
	Cmds Commands
}

// NewPkgManager returns Manager based on the OS plafrom
// It returns error when it's not able to inspect the host OS platform
// or the OS platform is not supported
func NewPkgManager() (Manager, error) {
	hostInfo, err := host.HostInfo()
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize package manager: %s", err)
	}
	switch hostInfo.PlatformFamily {
	case "debian":
		return NewAptManager()
	case "rhel":
		return nil, nil
	}

	return nil, nil
}
