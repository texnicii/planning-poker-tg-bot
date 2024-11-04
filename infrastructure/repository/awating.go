package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"planning_pocker_bot/domain/entity"
	"time"
)

type AwaitingRepository struct {
	db *gorm.DB
}

func (a *AwaitingRepository) Create(awaiting *entity.Awaiting) (*entity.Awaiting, error) {
	result := a.db.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"expiry_date"}),
	}).Create(awaiting)

	return awaiting, result.Error
}

func (a *AwaitingRepository) Get(chatId int64, callbackKey entity.CallbackKey) (*entity.Awaiting, error) {
	awaiting := &entity.Awaiting{}
	var result *gorm.DB
	if callbackKey == "" {
		result = a.db.Where("chat_id = ? AND expiry_date >= ?", chatId, time.Now()).
			Find(awaiting)
	} else {
		result = a.db.Where(
			"chat_id = ? AND callback_action_key = ? AND expiry_date >= ?",
			chatId,
			string(callbackKey),
			time.Now(),
		).Find(awaiting)
	}

	if awaiting.ChatId == 0 {
		return nil, result.Error
	}

	return awaiting, result.Error
}
