package infrastructure

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
	"planning_pocker_bot/infrastructure/repository"
	"planning_pocker_bot/infrastructure/telegram"
)

func Bootstrap() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("env loading fail: %s", errEnv)
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

	repository.RegisterRepositoriesAsServices(&appState)

	appState.Build()

	repository.MigrateSchema()
}

func TryEnv(envVar string, envDefault string) string {
	if val, ok := os.LookupEnv(envVar); ok {
		return val
	} else if envDefault == "" {
		log.Fatalf("%s is not defined", envVar)
	}

	return envDefault
}
