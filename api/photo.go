package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Photo entity definition
type Photo struct {
	ID          uint64 `gorm:"primary_key"`
	Name        string
	Description string
	Credit      string
	Place       string
	Like        int
	Country     Country
	CountryID   uint
	Tags        []Tag `gorm:"many2many:photo_tags;"`
	Comments    []Comment
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

// GetComments return all comments from a photo
func (s *Srv) GetComments(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var comments []Comment
	s.DB.Model(&Photo{ID: id}).Related(&comments)
	c.JSON(http.StatusOK, gin.H{"success": comments})
}

// GetTagsFromPhoto return all tags from a photo
func (s *Srv) GetTagsFromPhoto(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var tags []Tag
	s.DB.Model(&Photo{ID: id}).Related(&tags)
	c.JSON(http.StatusOK, gin.H{"success": tags})
}
