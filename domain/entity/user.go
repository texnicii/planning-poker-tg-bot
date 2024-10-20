package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `gorm:"type:varchar(100)"`
	Name     string `gorm:"type:varchar(100)"`
	AltName  string `gorm:"type:varchar(100)"`
}
