package migrate

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
)

func InitSchema() {
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
