// package manager implements service manager commands
package manager

import (
	"fmt"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/service"
	"github.com/milosgajdos83/servpeek/utils/service/commander"
)

// SvcManager provides service manager interface
type SvcManager interface {
	// Start starts service or returns error if the service can not be started
	Start(string) error
	// Stop stops service or returns error if the service can not be stopped
	Stop(string) error
	// Status return service status or returns error if the status can not be read
	Status(string) (service.Status, error)
}

// BaseSvcManager implements basic service manager commands
type BaseSvcManager struct {
	// cmd provides service commands
	cmd *commander.SvcCommander
}

// Start starts required service. It returns error if the service fails to start
func (bsm *BaseSvcManager) Start(svcName string) error {
	bsm.cmd.Start.Args = append([]string{svcName}, bsm.cmd.Start.Args...)
	_, err := bsm.cmd.Start.RunCombined()
	if err != nil {
		return err
	}
	return nil
}

// Stop stops required service. It returns error if the service fails to stop
func (bsm *BaseSvcManager) Stop(svcName string) error {
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
func (bsm *BaseSvcManager) Status(svcName string) (service.Status, error) {
	bsm.cmd.Status.Args = append(bsm.cmd.Status.Args, svcName)
	status, err := bsm.cmd.Stop.RunCombined()
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

// NewSvcManager returns SvcManager based on the service type
// It returns error if the SvcManager could not be created or required service type is not supported
func NewSvcManager(svcType string) (SvcManager, error) {
	switch svcType {
	case "init":
		return NewInitManager()
	case "upstart":
		return NewUpstartManager()
	case "systemd":
		return NewSystemdManager()
	}
	return nil, fmt.Errorf("Unsupported service type: %s", svcType)
}

// NewUpstartManager returns SvcManager or error
func NewInitManager() (SvcManager, error) {
	return &BaseSvcManager{
		cmd: commander.NewInitCommander(),
	}, nil
}

// NewUpstartManager returns SvcManager or error
func NewUpstartManager() (SvcManager, error) {
	return &BaseSvcManager{
		cmd: commander.NewUpstartCommander(),
	}, nil
}

// NewUpstartManager returns SvcManager or error
func NewSystemdManager() (SvcManager, error) {
	return &SystemdManager{}, nil
}
