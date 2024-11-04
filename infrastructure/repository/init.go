package repository

import (
	"gorm.io/gorm"
	"log"
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
)

func RegisterRepositoriesAsServices(app *di.Di) {
	app.Add(config.UserRepository, func() (any, error) {
		return &UserRepository{
			db: di.Get(config.DbClient).(*gorm.DB),
		}, nil
	}, 1)

	app.Add(config.AwaitingRepository, func() (any, error) {
		return &AwaitingRepository{
			db: di.Get(config.DbClient).(*gorm.DB),
		}, nil
	}, 1)
}

func MigrateSchema() {
	db := di.Get(config.DbClient).(*gorm.DB)
	entities := []any{
		&entity.User{},
		&entity.Awaiting{},
		&entity.Team{},
		&entity.UserTeam{},
	}

	for _, entitySchema := range entities {
		err := db.AutoMigrate(entitySchema)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
