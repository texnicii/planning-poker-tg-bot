package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ChatId   int64  `gorm:"primaryKey"`
	Nickname string `gorm:"type:varchar(100)"`
	AltName  string `gorm:"type:varchar(100)"`
}

func NewUser(chatId int64, nickname string, altName string) *User {
	return &User{ChatId: chatId, Nickname: nickname, AltName: altName}
}
