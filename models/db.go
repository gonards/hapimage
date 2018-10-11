package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)
// Skeleton is used to create all databases needed
func Skeleton() {
	db := openConnection()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Photo{}, &Country{}, &Tag{}, &Comment{}, &Card{})
}

func openConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
