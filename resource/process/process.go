// package process provides functions to perform various checks
// of various Linux process attributes
package process

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		ps := make([]*procfs.Proc, 0)
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

func processPrivileges(p *procfs.Proc, privRole string) ([]int, error) {
	// Process status path
	statusPath := fmt.Sprintf("%s/%d/status", procfs.DefaultMountPoint, p.PID)
	// open process status info file
	file, err := os.Open(statusPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Parse Real and Effective uid/gid
	var info string
	var realId, effId, sSet, fsId int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, privRole+":") {
			fmt.Sscanf(line, "%s %d %d %d %d", info, realId, effId, sSet, fsId)
			return []int{realId, effId}, nil
		}
	}
	// file scanner failed
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("Could not parse %v process status", p)
}

// IsCmdRunning checks if there is at least one process started with cmd command running
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

// IsPidRunning checks if there is a process with the given pid running
// It returns error if no such process is found on the host
func IsRunningPid(pid int) error {
	return withProcessPid(pid, func(p *procfs.Proc) error {
		return nil
	})
}

// IsRunningWithUid checks if all the procesess running given command
// are running it with provided effective Uid
// It returns error if either at least one process is running running with a different uid
// or no process with given cmd has been found
func IsRunningCmdWithUid(cmd string, username string) error {
	return withProcessCmd(cmd, func(procs []*procfs.Proc) error {
		uid, err := utils.RoleToId("user", username)
		if err != nil {
			return err
		}
		for _, proc := range procs {
			uids, err := processPrivileges(proc, "Uid")
			if err != nil {
				return err
			}
			if int(uid) != uids[1] {
				return fmt.Errorf("Incorrect process uid: %d found", uids[1])
			}
		}
		return nil
	})
}

// IsRunningWithGid checks if all the procesess running given command
// are running it with provided effective Gid
// It returns error if either at least one process is running running with a different gid
// or no process with given cmd has been found
func IsRunningCmdWithGid(cmd string, groupname string) error {
	return withProcessCmd(cmd, func(procs []*procfs.Proc) error {
		gid, err := utils.RoleToId("group", groupname)
		if err != nil {
			return err
		}
		for _, proc := range procs {
			gids, err := processPrivileges(proc, "Gid")
			if err != nil {
				return err
			}
			if int(gid) != gids[1] {
				return fmt.Errorf("Incorrect process gid: %d found", gids[1])
			}
		}
		return nil
	})
}
