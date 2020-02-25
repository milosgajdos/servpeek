package main

import "github.com/milosgajdos/servpeek/service"

// CheckServices checks various properties of OS services
// It returns error if any of the checked properties could not be satisfied.
func CheckServices() error {
	svc, err := service.NewSvc("docker", "upstart")
	if err != nil {
		return err
	}
	if err := service.IsRunning(svc); err != nil {
		return err
	}

	return nil
}
