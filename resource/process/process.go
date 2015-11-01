// Package process provides functions to query OS processes

// +build linux

package process

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils"
	"github.com/prometheus/procfs"
)

func withRunningProcs(fn func(procfs.Procs) error) error {
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
	return fn(procs)
}

func withProcessCmd(cmd string, fn func([]*procfs.Proc) error) error {
	return withRunningProcs(func(procs procfs.Procs) error {
		var ps []*procfs.Proc
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
		return fn(ps)
	})
}

func withProcessPid(pid int, fn func(*procfs.Proc) error) error {
	return withRunningProcs(func(procs procfs.Procs) error {
		for _, proc := range procs {
			if pid == proc.PID {
				return fn(&proc)
			}
		}
		return fmt.Errorf("%d process not found on the host", pid)
	})
}

func withRoleIDContext(role, name string, fn func(int) error) error {
	roleID, err := utils.RoleToID(role, name)
	if err != nil {
		return err
	}
	return fn(int(roleID))
}

func checkProcsPrivileges(procs []*procfs.Proc, id int, role string) error {
	for _, proc := range procs {
		// Process status path
		statusPath := fmt.Sprintf("%s/%d/status", procfs.DefaultMountPoint, proc.PID)
		// open process status info file
		file, err := os.Open(statusPath)
		if err != nil {
			return err
		}
		defer file.Close()
		// Parse Real and Effective uid/gid
		var info string
		var realID, effID, sSet, fsID int
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, role+":") {
				fmt.Sscanf(line, "%s %d %d %d %d", info, realID, effID, sSet, fsID)
			}
		}
		// file scanner failed
		if err := scanner.Err(); err != nil {
			return err
		}
		if id != effId {
			return fmt.Errorf("Expected PID: %d, Found PID: %d", id, proc.PID)
		}
	}
	return nil
}

// IsRunningCmd checks if there is at least one process started with cmd command running
// It return error if no such process could be found on the host
func IsRunningCmd(cmd string) error {
	return withProcessCmd(cmd, func(procs []*procfs.Proc) error {
		// ps slice is empty
		if len(procs) < 1 {
			return fmt.Errorf("Could not find %s process on the host", cmd)
		}
		return nil
	})
}

// IsRunningPid checks if there is a process with the given pid running
// It returns error if no such process is found on the host
func IsRunningPid(pid int) error {
	return withProcessPid(pid, func(p *procfs.Proc) error {
		return nil
	})
}

// IsRunningCmdWithUID checks if the provided command runs with provided user privileges
// It returns error if either provided process could not be found on the host or the
// process does not run with required user privileges
func IsRunningCmdWithUID(cmd string, username string) error {
	return withProcessCmd(cmd, func(procs []*procfs.Proc) error {
		return withRoleIDContext("User", username, func(id int) error {
			return checkProcsPrivileges(procs, id, "User")
		})
	})
}

// IsRunningCmdWithGID checks if the provided command runs with provided group privileges
// It returns error if either provided process could not be found on the host or the
// process does not run with required group privileges
func IsRunningCmdWithGID(cmd string, groupname string) error {
	return withProcessCmd(cmd, func(procs []*procfs.Proc) error {
		return withRoleIDContext("Group", groupname, func(id int) error {
			return checkProcsPrivileges(procs, id, "Group")
		})
	})
}

// ListRunning returns a slice of all running processes
// It returns an error of a process status could not be obtained
func ListRunning() ([]*resource.Process, error) {
	var ps []*resource.Process
	err := withRunningProcs(func(procs procfs.Procs) error {
		for _, proc := range procs {
			pstat, err := proc.NewStat()
			if err != nil {
				return err
			}
			ps = append(ps, &resource.Process{
				Pid: proc.PID,
				Cmd: pstat.Comm,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}
