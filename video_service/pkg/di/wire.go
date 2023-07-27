//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/stebinsabu13/microservice_video_streaming/video_service/pkg/Repository"
	"github.com/stebinsabu13/microservice_video_streaming/video_service/pkg/api"
	"github.com/stebinsabu13/microservice_video_streaming/video_service/pkg/api/service"
	"github.com/stebinsabu13/microservice_video_streaming/video_service/pkg/config"
	"github.com/stebinsabu13/microservice_video_streaming/video_service/pkg/db"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(db.Initdb, repository.NewVideoRepo, service.NewVideoServer, api.NewgrpcServe)
	return &api.Server{}, nil
}
