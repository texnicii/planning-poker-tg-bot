package planning_poker

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type NotSupported struct {
	handler.Model
}

func (ns NotSupported) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	response := new(messaging.ResponseBag)

	response.AddEditMessageResponseWithMarkup(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
		"🚫Planning poker is not support direct user chat\n↗️Add bot to some group",
		nil,
	)

	return response
}
