package service

// build linux
import "github.com/milosgajdos83/servpeek/utils/command"

// NewSysVCommander returns init service commander
func newSysVCommander() *Commander {
	return &Commander{
		StartCmd:  command.NewCommand(serviceCmd, []string{}...),
		StopCmd:   command.NewCommand(serviceCmd, []string{}...),
		StatusCmd: command.NewCommand(serviceCmd, []string{}...),
	}
}

// NewSysVInit returns SysInit or error
func NewSysVInit() SysInit {
	return &baseSysInit{
		Commander: newSysVCommander(),
		initType:  "sysv",
	}
}
