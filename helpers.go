package main

import "github.com/jinzhu/gorm"

func usernameExists(db *gorm.DB, username string) bool {
	usernameCount := 0
	db.Model(&User{}).Where("name = ?", username).Count(&usernameCount)
	return (usernameCount > 0)
}
