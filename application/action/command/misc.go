package cmd

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type UnknownCommandHandler struct {
	handler.Model
}

func (cmd UnknownCommandHandler) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	response := new(messaging.ResponseBag)
	if update.CallbackQuery != nil {
		// TODO i18n
		response.AddCallbackResponse(update.CallbackQuery.ID, "Hm... Not supported yet 🙄")
	} else {
		response.AddChatResponse(update.Message.Chat.ID, update.Message.MessageThreadID, "👀Unknown command")
	}

	return response
}

type Echo struct {
	*handler.Model
}

func (cmd Echo) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	response := new(messaging.ResponseBag)
	response.AddChatResponse(update.Message.Chat.ID, update.Message.MessageThreadID, "✅")

	return response
}
