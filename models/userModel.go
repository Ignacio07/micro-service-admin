package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Id       string `gorm:"not null" json:"id"`
	Store    string `gorm:"not null" json:"store"`
	Password string `gorm:"not null" json:"password"`
}
