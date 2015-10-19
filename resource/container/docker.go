package container

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsouza/go-dockerclient"
	"github.com/milosgajdos83/servpeek/resource"
)

// Reads DOCKER_HOST environment variable
// If it's empty, it validates and returns default docker endpoint
func getDockerEndPoint() (dep string, err error) {
	dockerHost := os.Getenv("DOCKER_HOST")
	if dockerHost != "" {
		dep = dockerHost
	} else {
		dep, err = docker.DefaultDockerHost()
		if err != nil {
			return
		}
	}
	return
}

// Reads DOCKER_CERT_PATH environment variable and returns a slice of strings
// that contains TLS ca, cert and key
func getDockerTLSCertPaths() []string {
	dockerCertPath := os.Getenv("DOCKER_CERT_PATH")
	if dockerCertPath != "" {
		ca := fmt.Sprintf("%s/ca.pem", dockerCertPath)
		cert := fmt.Sprintf("%s/cert.pem", dockerCertPath)
		key := fmt.Sprintf("%s/key.pem", dockerCertPath)
		return []string{ca, cert, key}
	}
	return nil
}

// Creates docker client and calls an anonymous func passed as argument
// If DOCKER_CERT_PATH is set, it creates HTTPS client, otherwise defaults to HTTP
func withDockerClient(fn func(*docker.Client) error) error {
	dockerEP, err := getDockerEndPoint()
	if err != nil {
		return err
	}
	dTLSCerPaths := getDockerTLSCertPaths()
	// Use HTTPS client
	if len(dTLSCerPaths) == 3 {
		ca, cert, key := dTLSCerPaths[0], dTLSCerPaths[1], dTLSCerPaths[2]
		client, err := docker.NewTLSClient(dockerEP, cert, key, ca)
		if err != nil {
			return err
		}
		return fn(client)
	}
	// Use plain HTTP client
	client, err := docker.NewClient(dockerEP)
	if err != nil {
		return err
	}
	return fn(client)
}

// Looks up *docker.APIImage which matches *resource.DockerImg passed as paramter and
// calls an anonymous func on it
func withDockerAPIImage(i *resource.DockerImg, fn func(*docker.APIImages) error) error {
	return withDockerClient(func(client *docker.Client) error {
		imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
		if err != nil {
			return err
		}
		for _, img := range imgs {
			for _, repoTag := range img.RepoTags {
				repo, tag := docker.ParseRepositoryTag(repoTag)
				if repo == i.Repo {
					if i.Tag == "" {
						if tag == "latest" {
							return fn(&img)
						}
					}
					if tag == i.Tag {
						return fn(&img)
					}
				}
			}
		}
		return fmt.Errorf("Image %s not found", i)
	})
}

// Looks up *docker.APIContainer that matches *resource.DockerContainer passed as parameter and
// calls an anonymous func on it
func withDockerAPIContainer(c *resource.DockerContainer, fn func(*docker.APIContainers) error) error {
	return withDockerClient(func(client *docker.Client) error {
		containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
		if err != nil {
			return err
		}
		for _, container := range containers {
			// Container ID has the highest priority
			if c.ID != "" && container.ID == c.ID {
				return fn(&container)
			}
			if c.Name != "" {
				if !strings.HasPrefix(c.Name, "/") {
					c.Name = "/" + c.Name
				}
				for _, name := range container.Names {
					if c.Name == name {
						return fn(&container)
					}
				}
			}
		}
		return fmt.Errorf("Container %s not found", c)
	})
}

// IsDockerImgPresent checks if the image is present on the Docker host
// It returns true if the images is present or error
func IsDockerImgPresent(img *resource.DockerImg) (bool, error) {
	err := withDockerAPIImage(img, func(apiImg *docker.APIImages) error {
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsDockerContainerPresent checks if docker container is present on the host
// It returns true if the images is present or error
func IsDockerContainerPresent(c *resource.DockerContainer) (bool, error) {
	err := withDockerAPIContainer(c, func(apiContainer *docker.APIContainers) error {
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
