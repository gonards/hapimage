package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Place entity definition
type Place struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string
	Photos    []Photo
	Country   Country
	CountryID uint64
}

// GetPlace return a place from specific id
func (s *Srv) GetPlace(c *gin.Context) {
	id := c.Param("id")
	var place Place
	if s.DB.First(&place, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Place not found " + id})
	} else {
		c.JSON(http.StatusOK, place)
	}
}

// PostPlace create a place
func (s *Srv) PostPlace(c *gin.Context) {
	var place Place
	if err := c.Bind(&place); err == nil {
		s.DB.Create(&place)
		c.JSON(http.StatusOK, gin.H{"success": place})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// GetPlaces return all places
func (s *Srv) GetPlaces(c *gin.Context) {
	var places []Place
	s.DB.Find(&places)
	c.JSON(http.StatusOK, gin.H{"success": places})
}

// GetPhotosFromPlace return all photos from a place
func (s *Srv) GetPhotosFromPlace(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var photos []Photo
	s.DB.Model(&Place{ID: id}).Related(&photos)
	c.JSON(http.StatusOK, gin.H{"success": photos})
}
