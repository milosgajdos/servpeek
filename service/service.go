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
	// SysInit returns instance of service sysinit which you can use
	// to control the service
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

// OsSvc is OS service
type OsSvc struct {
	// Name of the service
	name string
	// service manager
	sysInit SysInit
}

// NewOsSvc creates new Service or returns error if the requested service type is not supported
func NewOsSvc(name, sysInitType string) (*OsSvc, error) {
	sysInitTypes := []string{"upstart", "systemd", "sysv"}
	supported := make(map[string]bool)
	for _, sysInitType := range sysInitTypes {
		supported[sysInitType] = true
	}

	if !supported[sysInitType] {
		return nil, fmt.Errorf("Unsupported sysinit type: %s", sysInitType)
	}

	if name == "" {
		return nil, fmt.Errorf("Service name can not be empty")
	}

	s, err := NewSysInit(sysInitType)
	if err != nil {
		return nil, err
	}

	return &OsSvc{
		name:    name,
		sysInit: s,
	}, nil
}

// SysInit returns system init object that allows you to control os service
func (s *OsSvc) SysInit() SysInit {
	return s.sysInit
}

// Name returns name of the service
func (s *OsSvc) Name() string {
	return s.name
}

// String implements stringer interface
func (s *OsSvc) String() string {
	return fmt.Sprintf("[OsSvc] Name: %s, SysInit: %s", s.name, s.sysInit.Type())
}
