package service

import (
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
	_, err := repo.Create(data.ChatId, data.Nickname, data.AltName)
	if err != nil {
		return err
	}

	return nil
}
