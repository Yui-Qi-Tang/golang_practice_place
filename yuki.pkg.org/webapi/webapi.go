package webapi // this is package name

import (
	"github.com/gin-gonic/gin"
	// "os"
)

func RunGoGin() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/myFirstGoGin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello!!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
