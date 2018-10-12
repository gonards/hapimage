package api

import (
	"hapimage/models"

	"github.com/gin-gonic/gin"
)

// Cors middleware for api v1
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// NewServer create a server listening on port 8080
func NewServer() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/photo/:id", models.GetPhoto)
		v1.POST("/photo", models.PostPhoto)
		
		v1.GET("/card/:id", models.GetCard)
		v1.POST("/card", models.PostCard)
		
		v1.GET("/comment/:id", models.GetComment)
		v1.POST("/comment", models.PostComment)
		
		v1.GET("/country/:id", models.GetCountry)
		v1.POST("/country", models.PostCountry)
		
		v1.GET("/tag/:id", models.GetTag)
		v1.POST("/tag", models.PostTag)
	}
	r.Run(":8080")
}
