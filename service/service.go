package service

//build linux

import "fmt"

const (
	// service command
	serviceCmd = "service"
	// Running service status
	Running Status = iota + 1
	// Stopped service status
	Stopped
	// Unknown service status
	Unknown
)

// Service provides interface to interact with OS service
type Service interface {
	// Name returns the name of the service
	Name() string
	// SysInit returns the service init system
	SysInit() string
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

// OsOsSvc is OS service
type OsSvc struct {
	// Name of the service
	name string
	// Service Init system
	sysInit string
}

// NewOsSvc creates new Service or returns error if the requested service type is not supported
func NewOsSvc(name, sysInitType string) (*OsSvc, error) {
	sysInitTypes := []string{"upstart", "systemd", "sysv"}
	supported := make(map[string]bool)
	for _, sysInitType := range sysInitTypes {
		supported[sysInitType] = true
	}

	if !supported[sysInitType] {
		return nil, fmt.Errorf("Unsupported service init type: %s", sysInitType)
	}

	return &OsSvc{
		name:    name,
		sysInit: sysInitType,
	}, nil
}

// Name returns name of the service
func (s *OsSvc) Name() string {
	return s.name
}

// SysInit returns name of the sysinit type
func (s *OsSvc) SysInit() string {
	return s.sysInit
}

// String implements stringer interface
func (s *OsSvc) String() string {
	return fmt.Sprintf("[OsSvc] Name: %s, SysInit: %s", s.name, s.sysInit)
}
