package repository

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
)

func RegisterRepositoriesAsServices(app *di.Di) {
	app.Add(config.GroupRepository, func() (any, error) {
		return &GroupRepository{
			db: di.Get(config.DbClient).(*gorm.DB),
		}, nil
	}, 1)
}

func MigrateSchema() {
	db := di.Get(config.DbClient).(*gorm.DB)
	entities := []any{
		&entity.Group{},
		&entity.Vote{},
	}

	for _, entitySchema := range entities {
		err := db.AutoMigrate(entitySchema)
		if err != nil {
			log.Fatal().Msg(err.Error())
			return
		}
	}
}
