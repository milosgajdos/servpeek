package service

// build linux
import "github.com/milosgajdos83/servpeek/utils/command"

// SysVCommander provides SysV Init service manager commands
type SysVCommander struct {
	*SvcCommander
}

// NewSysVCommander returns init service commander
func NewSysVCommander() *SvcCommander {
	return &SvcCommander{
		Start:  command.NewCommand(serviceCmd, []string{"start"}...),
		Stop:   command.NewCommand(serviceCmd, []string{"stop"}...),
		Status: command.NewCommand(serviceCmd, []string{"status"}...),
	}
}

// NewSysVInit returns SysInit or error
func NewSysVInit() (SysInit, error) {
	return &BaseSysInit{
		cmd: NewSysVCommander(),
	}, nil
}
