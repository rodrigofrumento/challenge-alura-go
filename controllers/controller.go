package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigofrumento/challenge-alura-go/database"
	"github.com/rodrigofrumento/challenge-alura-go/models"
)

func ShowVideos(c *gin.Context) {
	var videos []models.Video
	database.DB.Find(&videos)
	c.JSON(200, videos)
}
