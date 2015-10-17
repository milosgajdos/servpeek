package container

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsouza/go-dockerclient"
	"github.com/milosgajdos83/servpeek/resource"
)

// creates docker client and calls whatever func with it
func withDockerClient(fn func(*docker.Client) error) error {
	var err error
	var dockerEndPoint string
	// Check if DOCKER_HOST env var is set
	dockerHost := os.Getenv("DOCKER_HOST")
	if dockerHost != "" {
		dockerEndPoint = dockerHost
	} else {
		dockerEndPoint, err = docker.DefaultDockerHost()
		if err != nil {
			return err
		}
	}
	// Check if DOCKER_CERT_PATH is set
	dockerCertPath := os.Getenv("DOCKER_CERT_PATH")
	if dockerCertPath != "" {
		ca := fmt.Sprintf("%s/ca.pem", dockerCertPath)
		cert := fmt.Sprintf("%s/cert.pem", dockerCertPath)
		key := fmt.Sprintf("%s/key.pem", dockerCertPath)
		client, err := docker.NewTLSClient(dockerEndPoint, cert, key, ca)
		if err != nil {
			return err
		}
		return fn(client)
	}
	client, err := docker.NewClient(dockerEndPoint)
	if err != nil {
		return err
	}
	return fn(client)
}

// DockerImgIsPresent checks if the image is present on the Docker host
func IsDockerImgPresent(img *resource.DockerImg) (bool, error) {
	err := withDockerClient(func(client *docker.Client) error {
		imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
		if err != nil {
			return err
		}
		for _, i := range imgs {
			for _, repoTag := range i.RepoTags {
				repo, tag := docker.ParseRepositoryTag(repoTag)
				if repo == img.Repo {
					// TODO: worth doing switch-case?
					if img.Tag == "" {
						if tag == "latest" {
							return nil
						}
					}
					if tag == img.Tag {
						return nil
					}
				}
			}
		}
		return fmt.Errorf("Image %s not found", img)
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsDockerContainerPresent checks if docker container is present on the host
func IsDockerContainerPresent(c *resource.DockerContainer) (bool, error) {
	err := withDockerClient(func(client *docker.Client) error {
		containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
		if err != nil {
			return err
		}
		for _, container := range containers {
			// Container ID has the highest priority
			if c.ID != "" && container.ID == c.ID {
				return nil
			}
			if c.Name != "" {
				if !strings.HasPrefix(c.Name, "/") {
					c.Name = "/" + c.Name
				}
				for _, name := range container.Names {
					if c.Name == name {
						return nil
					}
				}
			}
		}
		return fmt.Errorf("Container %s not found", c)
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
