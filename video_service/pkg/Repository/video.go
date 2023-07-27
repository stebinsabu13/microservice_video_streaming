package repository

import (
	"github.com/stebinsabu13/microservice_video_streaming/video_service/pkg/Repository/interfaces"
	"github.com/stebinsabu13/microservice_video_streaming/video_service/pkg/domain"
	"gorm.io/gorm"
)

type videoRepo struct {
	DB *gorm.DB
}

func NewVideoRepo(db *gorm.DB) interfaces.VideoRepo {
	return &videoRepo{
		DB: db,
	}
}

func (c *videoRepo) CreateVideoid(videoid string) error {
	if err := c.DB.Create(&domain.Video{VideoId: videoid}).Error; err != nil {
		return err
	}
	return nil
}
