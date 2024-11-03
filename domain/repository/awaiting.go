package repository

import (
	"planning_pocker_bot/domain/entity"
	"time"
)

// port to AwaitingRepository (implemented hexagonal design)

type AwaitingRepository interface {
	Create(chatId int64, waitKey entity.WaitKey, CommandCallback string, expired time.Time) (*entity.Awaiting, error)
	Get(chatId int64, waitKey entity.WaitKey) (*entity.Awaiting, error)
}
