// Package process provides functions to query OS processes
package process

import "fmt"

// Process provides simple API to query OS processes
type Process interface {
	// PID returns process id
	PID() int
	// Cmd returns that started the process
	Cmd() string
}

// OsProcess is a Linux OS process executed with command CMD
type OsProcess struct {
	// Process ID
	pid int
	// Process name
	cmd string
}

// NewOsProcess creates OsProcess object that can be used
func NewOsProcess(pid int, cmd string) (*OsProcess, error) {
	if pid <= 0 {
		return nil, fmt.Errorf("PID must be positivie integer: %d", pid)
	}

	if cmd == "" {
		return nil, fmt.Errorf("Process command can not be empty")
	}

	return &OsProcess{
		pid: pid,
		cmd: cmd,
	}, nil
}

// PID returns pid of the running process
func (p *OsProcess) PID() int { return p.pid }

// Cmd returns command that started the process
func (p *OsProcess) Cmd() string { return p.cmd }

// implement stringer interface
func (p *OsProcess) String() string {
	return fmt.Sprintf("[Process] PID: %d, Cmd: %s", p.pid, p.cmd)
}
