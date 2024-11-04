package entity

import (
	"time"
)

type CallbackKey string

type Awaiting struct {
	ChatId            int64       `gorm:"primaryKey;autoIncrement:false"`
	CallbackActionKey CallbackKey `gorm:"type:varchar(100)"`
	ExpiryDate        time.Time
}

func NewAwaiting(chatId int64, callbackKey CallbackKey, expiryDate time.Time) *Awaiting {
	return &Awaiting{ChatId: chatId, CallbackActionKey: callbackKey, ExpiryDate: expiryDate}
}
