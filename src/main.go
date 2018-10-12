package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(),types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Printf("%s\n", image.RepoTags[0])
	}

	containers, err := cli.ContainerList(context.Background(),
	 types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if !strings.Contains(container.Image, ":") {
			container.Image = container.Image + ":latest"
		}
		fmt.Printf("%s\n", container.Image)
	}
}
