package interfaces

import (
	"github.com/stebinsabu13/microservice_video_streaming/video_service/pkg/pb"
)

type VideoRepo interface {
	CreateVideoid(string) error
	FindAllVideo() ([]*pb.VideoID, error)
}
