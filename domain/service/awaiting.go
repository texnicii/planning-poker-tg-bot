package service

import (
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/domain/repository"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
	"time"
)

func HasAwaiting(chatId int64, key entity.WaitKey) bool {
	repo := di.Get(config.AwaitingRepository).(repository.AwaitingRepository)
	awaiting, _ := repo.Get(chatId, key)

	return awaiting != nil
}

func StartAwaiting(chatId int64, key entity.WaitKey, callback string, expired time.Time) error {
	repo := di.Get(config.AwaitingRepository).(repository.AwaitingRepository)
	_, err := repo.Create(chatId, key, callback, expired)
	if err != nil {
		return err
	}

	return nil
}

func GetAwaiting(chatId int64) (*entity.Awaiting, error) {
	repo := di.Get(config.AwaitingRepository).(repository.AwaitingRepository)

	return repo.Get(chatId, "")
}
