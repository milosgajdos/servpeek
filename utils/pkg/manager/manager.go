// package manager provides an implementation of package manager
package manager

// Manager defines generic software package manager interface
type Manager interface {
	// CheckInstalled returs true if pkgName package is installed
	CheckInstalled(pkgName string) (bool, error)
	// CheckInstalledVersion returns true if pkgName package is installed
	// and its version is pkgVersion
	CheckInstalledVersion(pkgName string, pkgVersion string) (bool, error)
}

// Commands defines various package manager commands
type Commands struct {
	// QueryInstalled queries installed packages
	QueryInstalled string
	// QueryInstalledVersion queries versions of installed packages
	QueryInstalledVersion string
}

// PkgManager is a package Manager that has a list of Commands
type PkgManager struct {
	// Cmds provides package manager commands
	cmds Commands
}

// NewPkgManager returns Manager based on the OS plafrom
// It returns error when it's not able to inspect the host OS platform
// or the OS platform is not supported
func NewPkgManager(pkgType string) (Manager, error) {
	switch pkgType {
	case "apt", "dpkg":
		return NewAptManager()
	case "rhel":
		return nil, nil
	}

	return nil, nil
}
