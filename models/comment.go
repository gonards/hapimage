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

// PostComment create a comment
func PostComment(c *gin.Context) {
	var comment Comment
	db := openConnection()
	defer db.Close()
	if err := c.Bind(&comment); err == nil {
		db.Create(&comment)
		c.JSON(http.StatusOK, gin.H{"success": comment})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}