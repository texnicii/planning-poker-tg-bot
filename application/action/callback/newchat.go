package callback

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type NewGame struct {
}

func (newChat NewGame) Handle(update tgbotapi.Update) messaging.ResponseBag {
	return messaging.ResponseBag{}
}
