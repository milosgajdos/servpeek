package service

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// SvcCommander provice SvcManager commands
type SvcCommander struct {
	// Start service
	Start *command.Command
	// Stop service
	Stop *command.Command
	// Check sercice status
	Status *command.Command
}

// NewSvcCommander returns SvcCommander or error if the required service typ is unsupported
func NewSvcCommander(sysInit string) (*SvcCommander, error) {
	switch sysInit {
	case "upstart":
		return NewUpstartCommander(), nil
	case "sysv":
		return NewSysVCommander(), nil
	}
	return nil, fmt.Errorf("Unsupported system init type: %s", sysInit)
}
