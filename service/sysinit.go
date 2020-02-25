package service

import (
	"fmt"
	"strings"

	"github.com/milosgajdos/servpeek/utils/command"
)

const (
	// ServiceCmd command
	ServiceCmd = "service"
	// SystemCtl command
	SystemCtl = "systemctl"
)

var svcStatusOut = map[string]map[string]string{
	"running": map[string]string{
		"sysv":    "running",
		"upstart": "running",
		"systemd": "active (running)",
	},
	"stopped": map[string]string{
		"sysv":    "is stopped",
		"upstart": "stop/waiting",
		"systemd": "inactive (stopped)",
	},
}

// SysInit allows to manage services of particular system init type
type SysInit interface {
	// Type returns type of system init
	Type() string
	// Start attempts to start a service of a given name
	// It returns error if the service could not be started
	Start(string) error
	// Stop attempts to stop a service of a given name
	// It returns error if the service could not be stopped
	Stop(string) error
	// Status returns service status or error if the status could not not be determined
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
	return nil, fmt.Errorf("Unsupported system init type: %s", sysInitType)
}

// Commander provides service management/control commands
type Commander struct {
	// Start service command
	StartCmd command.Command
	// Stop service command
	StopCmd command.Command
	// Check sercice status command
	StatusCmd command.Command
}

// baseSysInit provides basic service manager commands
type baseSysInit struct {
	// Service manager commands
	*Commander
	// system init type
	sysInitType string
}

// Type returns type of the system init
func (b *baseSysInit) Type() string {
	return b.sysInitType
}

// Start starts required service. It returns error if the service fails to start
func (b *baseSysInit) Start(svcName string) error {
	b.StartCmd.AppendArgs(sysInitCmdArgs(b.sysInitType, svcName, "start")...)
	_, err := b.StartCmd.RunCombined()
	return err
}

// Stop stops required service. It returns error if the service fails to stop
func (b *baseSysInit) Stop(svcName string) error {
	b.StopCmd.AppendArgs(sysInitCmdArgs(b.sysInitType, svcName, "stop")...)
	_, err := b.StopCmd.RunCombined()
	return err
}

// Status queries the status of service and returns it.
// It returns error if the service status could not be queried.
func (b *baseSysInit) Status(svcName string) (Status, error) {
	b.StatusCmd.AppendArgs(sysInitCmdArgs(b.sysInitType, svcName, "status")...)
	status, err := b.StatusCmd.RunCombined()
	if err != nil {
		return Unknown, err
	}
	switch {
	case strings.Contains(status, svcStatusOut["running"][b.sysInitType]):
		return Running, nil
	case strings.Contains(status, svcStatusOut["stopped"][b.sysInitType]):
		return Stopped, nil
	}
	return Unknown, fmt.Errorf("Unable to determine %s status", svcName)
}

func sysInitCmdArgs(syInitType, svcName, action string) []string {
	if syInitType == "systemd" {
		return []string{action, svcName + ".service"}
	}
	return []string{svcName, action}
}

// NewBaseCommander returns basic service commander
func NewBaseCommander(ctlCmd string) *Commander {
	return &Commander{
		StartCmd:  command.NewCommand(ctlCmd, []string{}...),
		StopCmd:   command.NewCommand(ctlCmd, []string{}...),
		StatusCmd: command.NewCommand(ctlCmd, []string{}...),
	}
}

// NewSysVInit returns SysInit which can manage SysV services
func NewSysVInit() SysInit {
	return &baseSysInit{
		Commander:   NewBaseCommander(ServiceCmd),
		sysInitType: "sysv",
	}
}

// NewUpstartInit returns SysInit which can manage upstart services
func NewUpstartInit() SysInit {
	return &baseSysInit{
		Commander:   NewBaseCommander(ServiceCmd),
		sysInitType: "upstart",
	}
}

// NewSystemdInit returns SysInit which can manage systemd services
func NewSystemdInit() SysInit {
	return &baseSysInit{
		Commander:   NewBaseCommander(SystemCtl),
		sysInitType: "systemd",
	}
}
