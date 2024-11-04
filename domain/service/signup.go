package service

import (
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/domain/repository"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
)

type SignUpDto struct {
	ChatId   int64
	Nickname string
	AltName  string
}

func SignUp(data SignUpDto) error {
	repo := di.Get(config.UserRepository).(repository.UserRepository)
	user, getErr := repo.Get(data.ChatId)
	if getErr != nil {
		return getErr
	}

	if user != nil {
		return nil
	}

	_, createErr := repo.Create(entity.NewUser(data.ChatId, data.Nickname, data.AltName))
	if createErr != nil {
		return createErr
	}

	return nil
}

func DeleteAccount(chatId int64) error {
	repo := di.Get(config.UserRepository).(repository.UserRepository)
	return repo.Delete(chatId)
}
