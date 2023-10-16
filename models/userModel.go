package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	
	Store    string `gorm:"not null" json:"store"`
	Password string `gorm:"not null" json:"password"`
}
