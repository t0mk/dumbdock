package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/samalba/dockerclient"
)

func main() {
	sck := os.Getenv("DOCKER_HOST")
	if sck == "" {
		sck = "unix:///var/run/docker.sock"
	}
	client, err := dockerclient.NewDockerClient(sck, nil)
	if err != nil {
		log.Fatal(err)
	}

	var showVolumes bool
	flag.BoolVar(&showVolumes, "v", false, "print volume per line")
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		log.Fatal("call properly")
	}

	cc, err := client.InspectContainer(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	for path, _ := range cc.Volumes {
		fmt.Println(path)
	}
}
