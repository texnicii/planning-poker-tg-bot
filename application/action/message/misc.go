package message

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type ChatMessageHandler struct {
	handler.Model
}

func (handler ChatMessageHandler) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	return nil
}
