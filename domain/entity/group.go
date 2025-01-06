package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Group struct {
	ChatId    int64     `gorm:"primaryKey;autoIncrement:false"`
	Options   Options   `gorm:"type:json;serializer:json"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Options struct {
	Lang string `json:"lang"`
}

func (o *Options) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *Options) Scan(src any) error {
	return json.Unmarshal(src.([]byte), o)
}
