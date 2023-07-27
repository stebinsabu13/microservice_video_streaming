package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/microservice_video_streaming/api_gateway/pkg/api/handler"
)

func RegisterVideoRoutes(api *gin.RouterGroup, videoHandler handler.VideoHandler) {
	api.POST("/upload", videoHandler.UploadVideo)
	api.GET("/stream/:video_id/:playlist", videoHandler.StreamVideo)
	api.GET("/video/all", videoHandler.FindAllVideo)
}
