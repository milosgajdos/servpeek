package resource

import "fmt"

// DockerImg defines Docker image
// It embeds pointer to docker.Image
type DockerImg struct {
	// Docker Repo
	Repo string
	// Docker image Tag
	Tag string
}

// Implement stringer interface
func (d *DockerImg) String() string {
	return fmt.Sprintf("[Docker Image] Name: %s, Tag: %s", d.Repo, d.Tag)
}

// DockerContainer defines Docker container
// It embeds pointer to docker container
type DockerContainer struct {
	// Docker container Id
	ID string
	// Docker container Name
	Name string
}

// Implement stringer interface
func (d *DockerContainer) String() (c string) {
	return fmt.Sprintf("[Docker Container] Name: %s, ID: %s", d.Name, d.ID)
}
