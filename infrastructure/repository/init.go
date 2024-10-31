package repository

import (
	"gorm.io/gorm"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
)

func RegisterRepositoriesAsServices(app *di.Di) {
	app.Add(config.UserRepository, func() (any, error) {
		return &UserRepository{
			db: di.Get(config.DbClient).(*gorm.DB),
		}, nil
	}, 1)
}
