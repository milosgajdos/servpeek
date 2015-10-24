package svc

import "fmt"

type Svc struct {
	// Name of the service
	Name string
	// Service manager type
	Type string
}

func (s *Svc) String() string {
	return fmt.Sprintf("[Service] Name: %s, Type: %s", s.Name, s.Type)
}
