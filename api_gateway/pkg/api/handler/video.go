package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/microservice_video_streaming/api_gateway/pkg/client/interfaces"
)

type VideoHandler struct {
	Client interfaces.VideoClient
}

func NewVideoHandler(client interfaces.VideoClient) VideoHandler {
	return VideoHandler{
		Client: client,
	}
}

func (cr *VideoHandler) UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to find the file",
			"error":   err.Error(),
		})
		return
	}
	res, err1 := cr.Client.UploadVideo(c.Request.Context(), file)
	if err1 != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to upload",
			"error":   err1.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"Success": res,
	})
}

func (cr *VideoHandler) StreamVideo(c *gin.Context) {
	filename := c.Param("video_id")
	playlist := c.Param("playlist")
	stream, err := cr.Client.StreamVideo(c.Request.Context(), filename, playlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to stream the video",
			"error":   err.Error(),
		})
		return
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "error while receiving video chunk:",
				"error":   err.Error(),
			})
			return
		}

		// Process each video chunk received from the server
		c.Header("Content-Type", "application/vnd.apple.mpegurl")
		c.Header("Content-Disposition", "inline")
		c.Writer.Write(resp.VideoChunk)
	}
}
