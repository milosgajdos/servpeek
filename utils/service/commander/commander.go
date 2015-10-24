// package commander implements service commander commands
package commander

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/utils/command"
)

const (
	serviceCmd = "service"
)

// SvcCommander provice SvcManager commands
type SvcCommander struct {
	// Check sercice status
	Status *command.Command
	// Start service
	Start *command.Command
	// Stop service
	Stop *command.Command
}

// NewSvcCommander returns SvcCommander or error if the required service typ is unsupported
func NewSvcCommander(svcType string) (*SvcCommander, error) {
	switch svcType {
	case "upstart":
		return NewUpstartCommander(), nil
	case "init":
		return NewInitCommander(), nil
	}
	return nil, fmt.Errorf("Unsupported service type: %s", svcType)
}
