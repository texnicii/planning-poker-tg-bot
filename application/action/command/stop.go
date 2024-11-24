package cmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/domain/service"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Stop struct {
	handler.Model
}

func (cmd Stop) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	_ = service.DeleteAccount(update.Message.Chat.ID)
	return nil
}
