package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Country entity definition
type Country struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

// GetCountry return a country from specific id
func GetCountry(c *gin.Context) {
	id := c.Param("id")
	var country Country
	db := openConnection()
	defer db.Close()
	if db.Debug().First(&country, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found " + id})
	} else {
		c.JSON(http.StatusOK, country)
	}
}

// PostCountry create a country
func PostCountry(c *gin.Context) {
	var country Country
	db := openConnection()
	defer db.Close()
	if err := c.Bind(&country); err == nil {
		db.Create(&country)
		c.JSON(http.StatusOK, gin.H{"success": country})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
