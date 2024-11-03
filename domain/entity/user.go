package entity

import "time"

type User struct {
	ChatId    int64     `gorm:"primaryKey;autoIncrement:false"`
	Nickname  string    `gorm:"type:varchar(100)"`
	AltName   string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func NewUser(chatId int64, nickname string, altName string) *User {
	return &User{ChatId: chatId, Nickname: nickname, AltName: altName}
}
