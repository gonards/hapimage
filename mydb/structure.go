package mydb

import (
	"github.com/jinzhu/gorm"
)

// Photo entity definition
type Photo struct {
	gorm.Model
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

// Country entity definition
type Country struct {
	gorm.Model
	Name string
}

// Tag entity definition
type Tag struct {
	gorm.Model
	Name   string
	Weight int
}

// Comment entity definition
type Comment struct {
	gorm.Model
	Username string
	Content  string
}

// Card entity definition
type Card struct {
	gorm.Model
	Type    string
	Photo   Photo
	PhotoID uint
	Caption string
	Link    string
}
