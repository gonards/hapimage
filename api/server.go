package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Srv server strcture implementation
type Srv struct {
	DB *gorm.DB
}

func (s *Srv) initDB(dbtype string) {
	switch dbtype {
	case "sqlite":
		s.DB = initDBSqlite()
	}
}

func (s *Srv) closeDB() {
	s.DB.Close()
}

// Cors middleware for api v1
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// NewServer create a server listening on port 8080
func NewServer() {
	s := Srv{}
	s.initDB("sqlite")
	defer s.closeDB()

	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/photo/:id", s.GetPhoto)
		v1.POST("/photo", s.PostPhoto)

		v1.GET("/card/:id", s.GetCard)
		v1.POST("/card", s.PostCard)

		v1.GET("/comment/:id", s.GetComment)
		v1.POST("/comment", s.PostComment)

		v1.GET("/country/:id", s.GetCountry)
		v1.POST("/country", s.PostCountry)

		v1.GET("/tag/:id", s.GetTag)
		v1.POST("/tag", s.PostTag)
	}
	r.Run(":8080")
}
