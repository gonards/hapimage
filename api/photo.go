package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Photo entity definition
type Photo struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
	Credit      string
	Place       string
	Like        int
	Country     Country
	CountryID   uint
	Tags        []Tag     `gorm:"many2many:photo_tags;"`
	Comments    []Comment `gorm:"many2many:photo_comment;"`
}

// GetPhoto return a photo from specific id
func (s *Srv) GetPhoto(c *gin.Context) {
	id := c.Param("id")
	var photo Photo
	s.DB.Where("ID = ?", id).First(&photo)
	c.JSON(http.StatusOK, photo)
}

// PostPhoto create a photo
func (s *Srv) PostPhoto(c *gin.Context) {
	var photo Photo
	if err := c.Bind(&photo); err == nil {
		s.DB.Create(&photo)
		c.JSON(http.StatusOK, gin.H{"success": photo})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}