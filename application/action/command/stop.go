package cmd

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"gorm.io/gorm"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Stop struct {
	handler.Model
	db *gorm.DB
}

func (cmd Stop) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	cmd.db.Delete(&entity.Group{}, update.Message.Chat.ID)

	return nil
}
