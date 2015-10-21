package resource

import "fmt"

// Process is a Linux OS process
type Process struct {
	// Process name
	Cmd string
}

// implement stringer interface
func (p *Process) String() string {
	return fmt.Sprintf("[Process] Name: %s", p.Cmd)
}

// Service is just a process
type Service struct {
	*Process
}

// implement stringer interface
func (s *Service) String() string {
	return fmt.Sprintf("{Service] Name: %s", s.Cmd)
}
