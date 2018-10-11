package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Card entity definition
type Card struct {
	ID      uint `gorm:"primary_key"`
	Type    string
	Photo   Photo
	PhotoID uint
	Caption string
	Link    string
}

// GetCard return a card from specific id
func GetCard(c *gin.Context) {
	id := c.Param("id")
	var card Card
	db := openConnection()
	defer db.Close()
	db.Where("ID = ?", id).First(&card)
	c.JSON(http.StatusOK, card)
}
