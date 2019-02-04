package main

import "github.com/jinzhu/gorm"

func usernameExists(db *gorm.DB, username string) bool {
	usernameCount := 0
	db.Model(&User{}).Where("name = ?", username).Count(&usernameCount)
	return (usernameCount > 0)
}

func removeSession(s []UserSession, i int) []UserSession {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
