package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Comment entity definition
type Comment struct {
	ID       uint `gorm:"primary_key"`
	Username string
	Content  string
}

// GetComment return a comment from specific id
func GetComment(c *gin.Context) {
	id := c.Param("id")
	var comment Comment
	db := openConnection()
	defer db.Close()
	db.Where("ID = ?", id).First(&comment)
	c.JSON(http.StatusOK, comment)
}
