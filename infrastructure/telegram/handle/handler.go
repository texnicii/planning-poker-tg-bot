package handle

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

const DefaultMessageHandlerAlias = "msg/default"

type HandlersContainer interface {
	Find(name string, isCallback bool) Handler
}

type Handler interface {
	SetInput(action string, input any)
	Handle(update tgbotapi.Update) *messaging.ResponseBag
}
