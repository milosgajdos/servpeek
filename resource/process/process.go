// package process provides functions to perform various checks
// of various Linux process attributes
package process

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/prometheus/procfs"
)

func withAllProcs(fn func(procfs.Procs) error) error {
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
	return withAllProcs(func(procs procfs.Procs) error {
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
	return withAllProcs(func(procs procfs.Procs) error {
		for _, proc := range procs {
			if pid == proc.PID {
				return fn(&proc)
			}
		}
		return fmt.Errorf("Could not find %d process on the host", pid)
	})
}

func processUidGid(p *procfs.Proc) (map[string][]int, error) {
	guidRE := regexp.MustCompile(`^[GU]id:.*`)
	procUidGidMap := make(map[string][]int)
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
		match := guidRE.FindStringSubmatch(line)
		if match == nil || len(match) < 5 {
			return nil, fmt.Errorf("Unable to parse process %v status", p)
		}
		fmt.Sscanf(line, "%s %d %d %d %d", info, realId, effId, sSet, fsId)
		procUidGidMap[strings.ToLower(info)] = []int{realId, effId}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return procUidGidMap, nil
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
func IsRunningCmdWithUid(cmd string) error {
	return withProcessCmd(cmd, func(procs []*procfs.Proc) error {
		return nil
	})
}

// IsRunningWithGid checks if all the procesess running given command
// are running it with provided effective Gid
// It returns error if either at least one process is running running with a different gid
// or no process with given cmd has been found
func IsRunningCmdWithGid(cmd string) error {
	return withProcessCmd(cmd, func(procs []*procfs.Proc) error {
		return nil
	})
}
