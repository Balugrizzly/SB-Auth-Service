package main

import (
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name        string `json:"Name"`
	Pw          string `json:"Pw"`
	IsSuperuser bool   `json:"IsSuperuser"`
}

// QUESTION: whats the auto update gorm code for unix time?

type UserSession struct {
	gorm.Model
	// this can be a bit confusing User is simultaneously a User Struct and telling gorm it belongs to User
	User          User
	UserID        uint
	SessionToken  string
	LoginTimeUnix int64
	LastSeenUnix  int64
}

// func (s *UserSession) updateSession(session UserSession) {
// 	s.SessionToken = session.SessionToken
// }
