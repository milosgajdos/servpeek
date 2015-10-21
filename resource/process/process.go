// package process provides functions to perform various checks
// of various Linux process attributes
package process

import (
	"fmt"

	"github.com/prometheus/procfs"
)

func withProcess(cmd string, fn func([]*procfs.Proc) error) error {
	ps := make([]*procfs.Proc, 0)
	// Mount /proc FS
	procFS, err := procfs.NewFS(procfs.DefaultMountPoint)
	if err != nil {
		return err
	}
	// List all running processes
	procs, err := procFS.AllProcs()
	if err != nil {
		return err
	}
	// Search for the ones that match passed command
	for _, proc := range procs {
		pstat, err := proc.NewStat()
		if err != nil {
			return err
		}
		if pstat.Comm == cmd {
			ps = append(ps, &proc)
		}
	}
	// ps slice is empty
	if len(ps) < 1 {
		return fmt.Errorf("Could not find %s process on the host", cmd)
	}
	return fn(ps)
}
