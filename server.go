package main

import (
	"net/http"

	"GINTEST/controller"
	"GINTEST/middlewares"
	"GINTEST/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	// router.LoadHTMLGlob("templates/*")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })
	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Findall())
	})

	server.POST("/Postvideos", func(c *gin.Context) {
		err := videoController.Save(c)
		print(err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Your Video Data is Valid!!"})
		}
	})

	server.Run()
}
