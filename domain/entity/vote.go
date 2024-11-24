package entity

import "time"

type Vote struct {
	GroupId   int64     `gorm:"primaryKey"`
	UserId    int64     `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`
	FirstName string    `gorm:"type:varchar(255)"`
	Icon      string    `gorm:"type:varchar(8)"`
	Value     string    `gorm:"type:varchar(32)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
