package repository

import (
	"gorm.io/gorm"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
)

func Read[T any](id any) *T {
	db := di.Get(config.DbClient).(*gorm.DB)
	entity := new(T)
	db.Find(&entity, id)

	return entity
}
