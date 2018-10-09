package mydb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Skeleton is used to create all databases needed
func Skeleton() {
	db, err := gorm.Open("mysql", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Photo{}, &Country{}, &Tag{}, &Comment{}, &Card{})
}
