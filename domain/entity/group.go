package entity

import "time"

type Group struct {
	ChatId    int64     `gorm:"primaryKey;autoIncrement:false"`
	Options   Options   `gorm:"type:json"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Options struct {
	Lang string
}
