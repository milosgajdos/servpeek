package resource

import "fmt"

// Svc provides OS service
type Svc struct {
	// Name of the service
	Name string
	// Service Init system
	SysInit string
}

// String implements stringer interface
func (s *Svc) String() string {
	return fmt.Sprintf("[Service] Name: %s, SysInit: %s", s.Name, s.SysInit)
}
