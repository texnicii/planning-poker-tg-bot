package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type ChatMessageHandler struct {
}

func (handler ChatMessageHandler) Handle(update tgbotapi.Update) messaging.ResponseBag {
	// nothing to reply
	return messaging.ResponseBag{}
}
