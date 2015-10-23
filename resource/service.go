package resource

import "fmt"

// Service is just a process
type Service struct {
	// Service name
	Name string
	// Type of service
	Type string
}

// implement stringer interface
func (s *Service) String() string {
	return fmt.Sprintf("{Service] Name: %s, Type: %s", s.Name, s.Type)
}
