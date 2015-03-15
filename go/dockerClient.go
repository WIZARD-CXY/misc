package main

// a rich api version of dockerclient
// reference http://godoc.org/github.com/fsouza/go-dockerclient
import (
	"fmt"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

//This interface abstract interface of dockerclient can do
type DockerInterface interface {
	ListContainers(options docker.ListContainersOptions) ([]docker.APIContainers, error)
	InspectContainer(id string) (*docker.Container, error)
	CreateContainer(docker.CreateContainerOptions) (*docker.Container, error)
	StartContainer(id string, hostConfig *docker.HostConfig) error
	StopContainer(id string, timeout uint) error
	RemoveContainer(opts docker.RemoveContainerOptions) error
	InspectImage(image string) (*docker.Image, error)
	ListImages(opts docker.ListImagesOptions) ([]docker.APIImages, error)
	PullImage(opts docker.PullImageOptions, auth docker.AuthConfiguration) error
	RemoveImage(image string) error
	Logs(opts docker.LogsOptions) error
	Version() (*docker.Env, error)
	CreateExec(docker.CreateExecOptions) (*docker.Exec, error)
	StartExec(string, docker.StartExecOptions) error
}

func main() {
	endpoint := os.Getenv("DH")

	if endpoint == "" {
		fmt.Println("DH env is not set and return ")
		return
	}
	fmt.Println(endpoint)
	client, err := docker.NewClient(endpoint)

	if err != nil {
		fmt.Println("client create failed")
	}
	/*imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentID)
	}*/

	//we only construct config here, docker just use config struct to create a container
	config := docker.CreateContainerOptions{
		//Name: "haha4",
		Config: &docker.Config{
			Hostname: "haha",
			Image:    "10.10.103.215:5000/sshd",
		},
	}

	dockerContainer, err := client.CreateContainer(config)

	if err != nil {
		fmt.Println("create error")
		return
	}

	fmt.Println("Created docker container with id", dockerContainer.ID)

	//we prepared a HostConfig used to start a container
	hc := &docker.HostConfig{
		NetworkMode: "none",
		Privileged:  true,
	}

	err = client.StartContainer(dockerContainer.ID, hc)

	if err != nil {
		fmt.Printf("Container started failed with ID %d\n", dockerContainer.ID)
	}
	fmt.Printf("Container started with ID %d\n", dockerContainer.ID)
}
