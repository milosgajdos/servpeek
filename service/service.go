package service

import "fmt"

const (
	// Running service status
	Running Status = iota + 1
	// Stopped service status
	Stopped
	// Unknown service status
	Unknown
)

// Service provides interface to interact with OS service
type Service interface {
	// Name returns the name of the OS service
	Name() string
	// Init returns service system init
	SysInit() SysInit
}

// Status defines service status
type Status int

// String implements stringer interface for Status
func (s Status) String() string {
	switch s {
	case Running:
		return "running"
	case Stopped:
		return "stopped"
	default:
		return "unknown"
	}
}

// Svc is OS service
type Svc struct {
	// Name of the service
	name string
	// System init type
	sysInit SysInit
}

// NewSvc creates new Service or returns error if the requested service type is not supported
func NewSvc(name, sysInitType string) (*Svc, error) {
	sysInit, err := NewSysInit(sysInitType)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, fmt.Errorf("Service name can not be empty")
	}

	return &Svc{
		name:    name,
		sysInit: sysInit,
	}, nil
}

// Name returns name of the service
func (s *Svc) Name() string {
	return s.name
}

// SysInit returns system init object that allows you to control os service
func (s *Svc) SysInit() SysInit {
	return s.sysInit
}

// String implements stringer interface
func (s *Svc) String() string {
	return fmt.Sprintf("[Svc] Name: %s, SysInit: %s", s.name, s.sysInit.Type())
}
