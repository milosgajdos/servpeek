package service

import "github.com/milosgajdos83/servpeek/utils/command"

// UpstartCommander provides upstart sys init commands
type UpstartCommander struct {
	*SvcCommander
}

// NewUpstartCommander returns upstart service commander
func NewUpstartCommander() *SvcCommander {
	return &SvcCommander{
		Start:  command.NewCommand(serviceCmd, []string{"start"}...),
		Stop:   command.NewCommand(serviceCmd, []string{"stop"}...),
		Status: command.NewCommand(serviceCmd, []string{"status"}...),
	}
}

// NewUpstartInit returns SysInit or error
func NewUpstartInit() (SysInit, error) {
	return &BaseSysInit{
		cmd: NewUpstartCommander(),
	}, nil
}
