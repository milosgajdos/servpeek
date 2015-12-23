package service

import "fmt"

// IsRunning checks if the supplied service is running
func IsRunning(s Service) error {
	// Check the service status
	status, err := s.SysInit().Status(s.Name())
	if err != nil {
		return err
	}
	// If the service isnt running, return error
	if status != Running {
		return fmt.Errorf("Service %s not running", s.Name())
	}
	return nil
}
