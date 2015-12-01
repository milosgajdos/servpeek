package container

import "fmt"

// DockerImg defines Docker image
// It has Repo and Tag
type DockerImg struct {
	// Docker Repo
	Repo string
	// Docker image Tag
	Tag string
}

// String implements stringer interface
func (d *DockerImg) String() string {
	return fmt.Sprintf("[Docker Image] Repo: %s, Tag: %s", d.Repo, d.Tag)
}

// DockerContainer defines Docker container
// It has ID and an Name
type DockerContainer struct {
	// Docker container Id
	ID string
	// Docker container Name
	Name string
}

// String implements stringer interface
func (d *DockerContainer) String() (c string) {
	return fmt.Sprintf("[Docker Container] ID: %s, Name: %s", d.ID, d.Name)
}
