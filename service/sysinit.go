package service

import (
	"fmt"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// SysInit provides service management commands
type SysInit interface {
	// Type returns type of system init to control services
	Type() string
	// Start starts a service or returns error if the service can not be started
	Start(string) error
	// Stop stops a service or returns error if the service can not be stopped
	Stop(string) error
	// Status returns service status or error if the status can not be determined
	Status(string) (Status, error)
}

// NewSysInit returns NewSysInit based on the system init type passed in as argument
// It returns error if the SysInit could not be created or required service type is not supported
func NewSysInit(sysInitType string) (SysInit, error) {
	switch sysInitType {
	case "sysv":
		return NewSysVInit(), nil
	case "upstart":
		return NewUpstartInit(), nil
	case "systemd":
		return NewSystemdInit(), nil
	}
	return nil, fmt.Errorf("Unsupported service init type: %s", sysInitType)
}

// Commander provides service management/control commands
type Commander struct {
	// Start service
	StartCmd command.Commander
	// Stop service
	StopCmd command.Commander
	// Check sercice status
	StatusCmd command.Commander
}

// baseSysInit provides basic service manager commands
type baseSysInit struct {
	// service manager
	*Commander
	// system init type
	initType string
}

// Type returns type of the system init
func (b *baseSysInit) Type() string {
	return b.initType
}

// Start starts required service. It returns error if the service fails to start
func (b *baseSysInit) Start(svcName string) error {
	b.StartCmd.AppendArgs(svcName, "start")
	_, err := b.StartCmd.RunCombined()
	return err
}

// Stop stops required service. It returns error if the service fails to stop
func (b *baseSysInit) Stop(svcName string) error {
	b.StopCmd.AppendArgs(svcName, "stop")
	_, err := b.StopCmd.RunCombined()
	return err
}

// Status queries the status of service and returns it.
// It returns error if the service status could not be queried.
// This method implements *SYSV INIT* and *UPSTART* status commands.
// You will have to override this method for other service managers
func (b *baseSysInit) Status(svcName string) (Status, error) {
	b.StatusCmd.AppendArgs(svcName, "status")
	status, err := b.StatusCmd.RunCombined()
	if err != nil {
		return Unknown, err
	}
	switch {
	case strings.Contains(status, "running"):
		return Running, nil
	case strings.Contains(status, "is stopped") || strings.Contains(status, "stop/waiting"):
		return Stopped, nil
	}
	return Unknown, fmt.Errorf("Unable to determine %s status", svcName)
}
