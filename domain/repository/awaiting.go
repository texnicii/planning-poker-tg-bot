package repository

import (
	"planning_pocker_bot/domain/entity"
)

// port to AwaitingRepository (implemented hexagonal design)

type AwaitingRepository interface {
	Create(awaiting *entity.Awaiting) (*entity.Awaiting, error)
	Get(chatId int64, callbackKey entity.CallbackKey) (*entity.Awaiting, error)
}
