package api

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
func (s *Srv) GetComment(c *gin.Context) {
	id := c.Param("id")
	var comment Comment
	s.DB.Where("ID = ?", id).First(&comment)
	c.JSON(http.StatusOK, comment)
}

// PostComment create a comment
func (s *Srv) PostComment(c *gin.Context) {
	var comment Comment
	if err := c.Bind(&comment); err == nil {
		s.DB.Create(&comment)
		c.JSON(http.StatusOK, gin.H{"success": comment})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}