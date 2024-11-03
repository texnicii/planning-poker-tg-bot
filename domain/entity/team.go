package entity

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)"`
	OwnerChatId int64  `gorm:"index"`
}
