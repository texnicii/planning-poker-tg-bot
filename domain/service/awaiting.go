package service

import (
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/domain/repository"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
	"time"
)

func StartAwaiting(chatId int64, callbackKey entity.CallbackKey, expired time.Time) error {
	repo := di.Get(config.AwaitingRepository).(repository.AwaitingRepository)
	_, err := repo.Create(entity.NewAwaiting(chatId, callbackKey, expired))
	if err != nil {
		return err
	}

	return nil
}

func GetAwaiting(chatId int64) (*entity.Awaiting, error) {
	repo := di.Get(config.AwaitingRepository).(repository.AwaitingRepository)

	return repo.Get(chatId, "")
}
