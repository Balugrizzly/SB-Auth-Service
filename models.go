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

type UserSession struct {
	gorm.Model
	User          User
	SessionToken  string
	LoginTimeUnix int64
	LastSeenUnix  int64
}
