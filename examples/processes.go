package main

import "github.com/milosgajdos/servpeek/process"

// CheckProcesses checks various properties of os processes
// It returns error if any of the checked properties could not be satisfied.
func CheckProcesses() error {
	if err := process.IsRunningCmd("docker"); err != nil {
		return err
	}

	if err := process.IsRunningCmdWithUID("docker", "root"); err != nil {
		return err
	}

	return nil
}
