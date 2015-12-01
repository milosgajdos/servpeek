package service

import "github.com/milosgajdos83/servpeek/utils/command"

// NewUpstartCommander returns upstart service commander
func newUpstartCommander() *Commander {
	return &Commander{
		StartCmd:  command.NewCommand(serviceCmd, []string{}...),
		StopCmd:   command.NewCommand(serviceCmd, []string{}...),
		StatusCmd: command.NewCommand(serviceCmd, []string{}...),
	}
}

// NewUpstartInit returns SysInit or error
func NewUpstartInit() SysInit {
	return &baseSysInit{
		Commander: newUpstartCommander(),
		initType:  "upstart",
	}
}
