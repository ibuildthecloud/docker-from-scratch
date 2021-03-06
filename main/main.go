package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	dockerlaunch "github.com/rancher/docker-from-scratch"
)

func main() {
	if os.Getenv("DOCKER_LAUNCH_DEBUG") == "true" {
		log.SetLevel(log.DebugLevel)
	}

	if len(os.Args) < 2 {
		log.Fatalf("Usage Example: %s /usr/bin/docker -d -D", os.Args[0])
	}

	args := []string{}
	if len(os.Args) > 1 {
		args = os.Args[2:]
	}

	err := dockerlaunch.LaunchDocker(os.Args[1], args...)
	if err != nil {
		log.Fatal(err)
	}
}
