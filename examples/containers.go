package main

import "github.com/milosgajdos83/servpeek/container"

// CheckContainers checkss various properties of container images and containers
// It returns error if any of the checked properties could not be satisfied.
func CheckContainers() error {
	dockerImg := &container.DockerImg{
		Repo: "busybox",
	}
	if err := container.IsDockerImgPresent(dockerImg); err != nil {
		return err
	}

	dockerContainer := &container.DockerContainer{
		Name: "pensive_mahavira",
	}
	if err := container.IsDockerContainerPresent(dockerContainer); err != nil {
		return err
	}

	return nil
}
