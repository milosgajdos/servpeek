// package init implements service manager commands
package sysinit

import (
	"fmt"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/service"
	"github.com/milosgajdos83/servpeek/utils/service/commander"
)

// SvcInit provides service management commands
type SvcInit interface {
	// Start starts service or returns error if the service can not be started
	Start(string) error
	// Stop stops service or returns error if the service can not be stopped
	Stop(string) error
	// Status return service status or returns error if the status can not be read
	Status(string) (service.Status, error)
}

// BaseSvcInit implements basic service manager commands
type BaseSvcInit struct {
	// cmd provides service commands
	cmd *commander.SvcCommander
}

// Start starts required service. It returns error if the service fails to start
func (bsm *BaseSvcInit) Start(svcName string) error {
	bsm.cmd.Start.Args = append([]string{svcName}, bsm.cmd.Start.Args...)
	_, err := bsm.cmd.Start.RunCombined()
	if err != nil {
		return err
	}
	return nil
}

// Stop stops required service. It returns error if the service fails to stop
func (bsm *BaseSvcInit) Stop(svcName string) error {
	bsm.cmd.Stop.Args = append([]string{svcName}, bsm.cmd.Stop.Args...)
	_, err := bsm.cmd.Stop.RunCombined()
	if err != nil {
		return err
	}
	return nil
}

// Status queries the status of service and returns it.
// It returns error if the service status could not be queried.
// This method implements *SYSV INIT* and *UPSTART* status commands.
// You will have to override this method for other service managers
func (bsm *BaseSvcInit) Status(svcName string) (service.Status, error) {
	bsm.cmd.Status.Args = append([]string{svcName}, bsm.cmd.Status.Args...)
	status, err := bsm.cmd.Status.RunCombined()
	if err != nil {
		return service.Unknown, err
	}
	switch {
	case strings.Contains(status, "is running") || strings.Contains(status, "start/running"):
		return service.Running, nil
	case strings.Contains(status, "is stopped") || strings.Contains(status, "stop/waiting"):
		return service.Stopped, nil
	}
	return service.Unknown, fmt.Errorf("Unable to determine %s status", svcName)
}

// NewSvcInit returns NewSvcInit based on the system init type passed in as argument
// It returns error if the SvcInit could not be created or required service type is not supported
func NewSvcInit(sysInit string) (SvcInit, error) {
	switch sysInit {
	case "sysv":
		return NewSysVInit()
	case "upstart":
		return NewUpstartInit()
	case "systemd":
		return NewSystemdInit()
	}
	return nil, fmt.Errorf("Unsupported service type: %s", sysInit)
}

// NewUpstartInit returns SvcInit or error
func NewSysVInit() (SvcInit, error) {
	return &BaseSvcInit{
		cmd: commander.NewSysVCommander(),
	}, nil
}

// NewUpstartInit returns SvcInit or error
func NewUpstartInit() (SvcInit, error) {
	return &BaseSvcInit{
		cmd: commander.NewUpstartCommander(),
	}, nil
}

// NewSystemdInit returns SvcInit or error
func NewSystemdInit() (SvcInit, error) {
	return &SystemdInit{}, nil
}
