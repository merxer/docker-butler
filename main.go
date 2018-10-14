package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Filter(s []string, r string) []string {
	filterString := make([]string, 0)
	for _, v := range s {
		if v != r {
			filterString = append(filterString, v)
		}
	}
	return filterString
}

func main() {
	allImages := make([]string,0)
	unusedImages := make([]string,0)

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(),types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		allImages = append(allImages,image.RepoTags[0])
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	if len(containers) == 0 { unusedImages = allImages }

	for _, container := range containers {
		if !strings.Contains(container.Image, ":") {
			container.Image = container.Image + ":latest"
		}
		unusedImages = Filter(allImages, container.Image)
	}

	for _, removeImage := range unusedImages {
		ok, err := cli.ImageRemove(context.Background(),
		                           removeImage,types.ImageRemoveOptions{Force: true, PruneChildren: true})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", ok)
	}
}
