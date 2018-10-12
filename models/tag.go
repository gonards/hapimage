package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tag entity definition
type Tag struct {
	ID     uint   `gorm:"primary_key"`
	Name   string `gorm:"UNIQUE"`
	Weight int    `gorm:"default:1"`
}

// GetTag return a tag from specific id
func GetTag(c *gin.Context) {
	id := c.Param("id")
	var tag Tag
	db := openConnection()
	defer db.Close()
	db.Where("ID = ?", id).First(&tag)
	c.JSON(http.StatusOK, tag)
}

// PostTag create a tag
func PostTag(c *gin.Context) {
	var tag Tag
	db := openConnection()
	defer db.Close()
	if err := c.Bind(&tag); err == nil {
		result := db.Where("name = ?", tag.Name).First(&tag)
		if !result.RecordNotFound() {
			db.Model(&tag).Update("weight", tag.Weight+1)
		} else {
			db.Create(&tag)
		}
		c.JSON(http.StatusOK, gin.H{"success": tag})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
