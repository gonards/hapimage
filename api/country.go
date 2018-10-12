package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Country entity definition
type Country struct {
	ID     uint64 `gorm:"primary_key"`
	Name   string
	Photos []Photo
}

// GetCountry return a country from specific id
func (s *Srv) GetCountry(c *gin.Context) {
	id := c.Param("id")
	var country Country
	if s.DB.First(&country, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found " + id})
	} else {
		c.JSON(http.StatusOK, country)
	}
}

// PostCountry create a country
func (s *Srv) PostCountry(c *gin.Context) {
	var country Country
	if err := c.Bind(&country); err == nil {
		s.DB.Create(&country)
		c.JSON(http.StatusOK, gin.H{"success": country})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
