package models

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
func GetPhoto(c *gin.Context) {
	id := c.Param("id")
	var photo Photo
	db := openConnection()
	defer db.Close()
	db.Where("ID = ?", id).First(&photo)
	c.JSON(http.StatusOK, photo)
}

// PostPhoto create a photo
func PostPhoto(c *gin.Context) {
	var photo Photo
	db := openConnection()
	defer db.Close()
	if err := c.Bind(&photo); err == nil {
		db.Create(&photo)
		c.JSON(http.StatusOK, gin.H{"success": photo})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}