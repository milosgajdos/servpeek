package service

import (
	"fmt"
	"strings"
)

// SysInit provides service management commands
type SysInit interface {
	// Start starts a service or returns error if the service can not be started
	Start(string) error
	// Stop stops a service or returns error if the service can not be stopped
	Stop(string) error
	// Status returns service status or error if the status can not be determined
	Status(string) (Status, error)
}

// NewSysInit returns NewSysInit based on the system init type passed in as argument
// It returns error if the SysInit could not be created or required service type is not supported
func NewSysInit(sysInit string) (SysInit, error) {
	switch sysInit {
	case "sysv":
		return NewSysVInit()
	case "upstart":
		return NewUpstartInit()
	case "systemd":
		return NewSystemdInit()
	}
	return nil, fmt.Errorf("Unsupported service init type: %s", sysInit)
}

// BaseSysInit provides basic service manager commands
type BaseSysInit struct {
	// cmd provides service commands
	cmd *SvcCommander
}

// Start starts required service. It returns error if the service fails to start
func (b *BaseSysInit) Start(svcName string) error {
	b.cmd.Start.Args = append([]string{svcName}, b.cmd.Start.Args...)
	_, err := b.cmd.Start.RunCombined()
	return err
}

// Stop stops required service. It returns error if the service fails to stop
func (b *BaseSysInit) Stop(svcName string) error {
	b.cmd.Stop.Args = append([]string{svcName}, b.cmd.Stop.Args...)
	_, err := b.cmd.Stop.RunCombined()
	return err
}

// Status queries the status of service and returns it.
// It returns error if the service status could not be queried.
// This method implements *SYSV INIT* and *UPSTART* status commands.
// You will have to override this method for other service managers
func (b *BaseSysInit) Status(svcName string) (Status, error) {
	b.cmd.Status.Args = append([]string{svcName}, b.cmd.Status.Args...)
	status, err := b.cmd.Status.RunCombined()
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
