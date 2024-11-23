package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	// Initialize Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// Define the container name to remove
	containerName := "happy_shtern"

	// List all containers (including stopped ones)
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}

	// Find the container by name
	var containerID string
	for _, c := range containers {
		for _, name := range c.Names {
			if name == "/"+containerName { // Container names have a leading "/"
				containerID = c.ID
				break
			}
		}
		if containerID != "" {
			break
		}
	}

	// Check if the container was found
	if containerID == "" {
		fmt.Printf("Container with name %s not found\n", containerName)
		return
	}

	fmt.Printf("Removing container with ID: %s\n", containerID)

	// Remove the container
	err = cli.ContainerRemove(context.Background(), containerID, container.RemoveOptions{
		Force:         true, // Force remove if running
		RemoveVolumes: true, // Remove associated volumes
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Container removed successfully")
}
