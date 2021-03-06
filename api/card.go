package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Card entity definition
type Card struct {
	ID      uint64 `gorm:"primary_key"`
	Type    string
	Photo   Photo
	PhotoID uint
	Caption string
	Link    string
	Country Country
}

// GetCard return a card from specific id
func (s *Srv) GetCard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var card Card
	s.DB.Where("ID = ?", id).First(&card)
	c.JSON(http.StatusOK, card)
}

// PostCard create a card
func (s *Srv) PostCard(c *gin.Context) {
	var card Card
	if err := c.Bind(&card); err == nil {
		s.DB.Create(&card)
		c.JSON(http.StatusOK, gin.H{"success": card})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// GetCards return all Cards
func (s *Srv) GetCards(c *gin.Context) {
	var cards []Card
	s.DB.Find(&cards)
	c.JSON(http.StatusOK, gin.H{"success": cards})
}

// GetLastCards return last nb cards
func (s *Srv) GetLastCards(c *gin.Context) {
	nb, err := strconv.ParseUint(c.Param("nb"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var cards []Card
	s.DB.Limit(nb).Order("id desc").Find(&cards)
	c.JSON(http.StatusOK, gin.H{"success": cards})
}
