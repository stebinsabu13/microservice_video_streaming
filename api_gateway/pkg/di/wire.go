//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/stebinsabu13/microservice_video_streaming/api_gateway/pkg/api"
	"github.com/stebinsabu13/microservice_video_streaming/api_gateway/pkg/api/handler"
	"github.com/stebinsabu13/microservice_video_streaming/api_gateway/pkg/client"
	"github.com/stebinsabu13/microservice_video_streaming/api_gateway/pkg/config"
)

func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(client.InitClient, client.NewVideoClient, handler.NewVideoHandler, api.NewServeHTTP)
	return &api.Server{}, nil
}
