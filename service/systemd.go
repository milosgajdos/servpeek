package service

// build linux

import (
	"fmt"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const (
	systemctl = "systemctl"
)

///////////////////////////////////
// TODO: re-implement using dbus //
///////////////////////////////////

// systemdInit provides systemd init commands
type systemdInit struct{}

// NewSystemdInit returns SysInit or error
func NewSystemdInit() SysInit {
	return &systemdInit{}
}

// Type returns type of system init to control services
func (s *systemdInit) Type() string {
	return "systemd"
}

// Start starts systemd service. It returns error if the service fails to be started
func (s *systemdInit) Start(svcName string) error {
	startCmd := command.NewCommand(systemctl, []string{"start", svcName + ".service"}...)
	_, err := startCmd.RunCombined()
	if err != nil {
		return err
	}
	return nil
}

// Stop stops systemd service. It returns error if the service fails to be stopped
func (s *systemdInit) Stop(svcName string) error {
	stopCmd := command.NewCommand(systemctl, []string{"stop", svcName + ".service"}...)
	_, err := stopCmd.RunCombined()
	if err != nil {
		return err
	}
	return nil
}

// Status queries status of systemd service and returns it.
// It returns error if the status of the queried service could not be determined
func (s *systemdInit) Status(svcName string) (Status, error) {
	statusCmd := command.NewCommand(systemctl, []string{"status", svcName + ".service"}...)
	status, err := statusCmd.RunCombined()
	if err != nil {
		return Unknown, err
	}
	switch {
	case strings.Contains(status, "active (running)"):
		return Running, nil
	case strings.Contains(status, "inactive (stopped)"):
		return Stopped, nil
	}
	return Unknown, fmt.Errorf("Unable to determine %s status", svcName)
}
