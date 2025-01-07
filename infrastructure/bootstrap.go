package infrastructure

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
	"planning_pocker_bot/infrastructure/migrate"
	"planning_pocker_bot/infrastructure/telegram"
)

func Bootstrap() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Warn().Msg(errEnv.Error())
	}

	// init services and add to global service container
	appState := di.NewApp()
	appState.Add(config.BotClient, func() (any, error) {
		tgBot, err := telegram.NewBotClient(TryEnv("TG_API_TOKEN", ""))
		return &tgBot, err
	}, 0)

	appState.Add(config.DbClient, func() (any, error) {
		db, err := gorm.Open(mysql.Open(TryEnv("DSN", "")), &gorm.Config{})
		return db, err
	}, 99)

	appState.Build()

	migrate.InitSchema()
}

func TryEnv(envVar string, envDefault string) string {
	if val, ok := os.LookupEnv(envVar); ok {
		return val
	} else if envDefault == "" {
		log.Fatal().Msgf("%s is not defined", envVar)
	}

	return envDefault
}
