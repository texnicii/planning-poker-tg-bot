package repository

import (
	"planning_pocker_bot/domain/entity"
)

// port to UserRepository (implemented hexagonal design)

type UserRepository interface {
	Create(chatId int64, nickname string, altName string) (*entity.User, error)
	Delete(chatId int64) error
	Get(chatId int64) (*entity.User, error)
}
