package repository

import (
	"planning_pocker_bot/domain/entity"
)

// port to GroupRepository (implemented hexagonal design)

type GroupRepository interface {
	Create(user *entity.Group) (*entity.Group, error)
	Delete(chatId int64) error
	Get(chatId int64) (*entity.Group, error)
}
