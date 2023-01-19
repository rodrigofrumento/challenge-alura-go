package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigofrumento/challenge-alura-go/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/videos", controllers.ShowVideos)
	r.Run()
}
