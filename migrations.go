package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func autoMigration(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&User{})
}
