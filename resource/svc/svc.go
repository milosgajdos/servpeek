package svc

//build linux

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/service"
	"github.com/milosgajdos83/servpeek/utils/service/sysinit"
)

// IsRunning checks if the service is running
func IsRunning(svc *resource.Svc) error {
	svcInit, err := sysinit.NewSvcInit(svc.SysInit)
	if err != nil {
		return err
	}
	// Check the service status
	status, err := svcInit.Status(svc.Name)
	if err != nil {
		return err
	}
	// If the service isnt running, return error
	if status != service.Running {
		return fmt.Errorf("Service %s not running", svc.Name)
	}
	return nil
}
