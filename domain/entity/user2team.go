package entity

type UserTeam struct {
	TeamId   uint  `gorm:"primaryKey;autoIncrement:false"`
	ChatId   int64 `gorm:"primaryKey;autoIncrement:false"`
	IsActive bool
}
