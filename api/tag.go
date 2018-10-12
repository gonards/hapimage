package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tag entity definition
type Tag struct {
	ID     uint64   `gorm:"primary_key"`
	Name   string `gorm:"UNIQUE"`
	Weight int    `gorm:"default:1"`
}

// GetTag return a tag from specific id
func (s *Srv) GetTag(c *gin.Context) {
	id := c.Param("id")
	var tag Tag
	s.DB.Where("ID = ?", id).First(&tag)
	c.JSON(http.StatusOK, tag)
}

// PostTag create a tag
func (s *Srv) PostTag(c *gin.Context) {
	var tag Tag
	if err := c.Bind(&tag); err == nil {
		result := s.DB.Where("name = ?", tag.Name).First(&tag)
		if !result.RecordNotFound() {
			s.DB.Model(&tag).Update("weight", tag.Weight+1)
		} else {
			s.DB.Create(&tag)
		}
		c.JSON(http.StatusOK, gin.H{"success": tag})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// GetPhotosFromTags return all photos from a tag
func (s *Srv) GetPhotosFromTags(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var photos []Photo
	s.DB.Model(&Tag{ID: id}).Related(&photos)
	c.JSON(http.StatusOK, gin.H{"success": photos})
}
