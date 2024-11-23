package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.41"), client.FromEnv)
	if err != nil {
		panic(err)
	}

	fmt.Println("Docker API Version:", cli.ClientVersion())

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		if len(ctr.Names) > 0 {
			name := ctr.Names[0][1:] // Take the first name and remove the leading '/'
			fmt.Printf("%s %s\n", name, ctr.Image)
		}

	}

}
