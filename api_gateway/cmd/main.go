package main

import (
	"log"

	"github.com/stebinsabu13/microservice_video_streaming/api_gateway/pkg/config"
	"github.com/stebinsabu13/microservice_video_streaming/api_gateway/pkg/di"
)

func main() {
	c, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config:", configerr)
	}
	server, dierr := di.InitializeAPI(c)
	if dierr != nil {
		log.Fatal("cannot initialize server", dierr)
	}
	server.Start()
}
