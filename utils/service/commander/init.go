package commander

import "github.com/milosgajdos83/servpeek/utils/command"

// InitCommander provides SysV Init service manager commands
type InitCommander struct {
	*SvcCommander
}

// NewInitCommander returns init service commander
func NewInitCommander() *SvcCommander {
	return &SvcCommander{
		Start:  command.NewCommand(serviceCmd, []string{"start"}...),
		Stop:   command.NewCommand(serviceCmd, []string{"stop"}...),
		Status: command.NewCommand(serviceCmd, []string{"status"}...),
	}
}
