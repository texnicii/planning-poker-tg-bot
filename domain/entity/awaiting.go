package entity

import (
	"time"
)

type WaitKey string

type Awaiting struct {
	ChatId          int64   `gorm:"primaryKey;autoIncrement:false"`
	WaitKey         WaitKey `gorm:"type:varchar(100)"`
	CommandCallback string  `gorm:"type:varchar(100)"`
	ExpiryDate      time.Time
}

func NewAwaiting(chatId int64, waitKey WaitKey, commandCallback string, expiryDate time.Time) *Awaiting {
	return &Awaiting{ChatId: chatId, WaitKey: waitKey, CommandCallback: commandCallback, ExpiryDate: expiryDate}
}
