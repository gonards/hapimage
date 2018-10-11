package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tag entity definition
type Tag struct {
	ID     uint `gorm:"primary_key"`
	Name   string
	Weight int
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
