package api

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitSkeleton is used to create all databases needed
func InitSkeleton(sqlType string) {
	var db *gorm.DB
	switch sqlType {
	case "sqlite":
		db = initDBSqlite()
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Photo{}, &Country{}, &Tag{}, &Comment{}, &Card{})
}

// initDBSqlite open connection to sqlite db 
func initDBSqlite() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalf("failed to connect database")
	}
	return db
}
