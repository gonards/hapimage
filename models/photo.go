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
