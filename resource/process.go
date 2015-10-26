package resource

import "fmt"

// Process is a Linux OS process executed with command CMD
type Process struct {
	// Process ID
	Pid int
	// Process name
	Cmd string
}

// implement stringer interface
func (p *Process) String() string {
	return fmt.Sprintf("[Process] PID: %d, Cmd: %s", p.Cmd, p.Pid)
}
