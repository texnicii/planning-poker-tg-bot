package handle

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
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
