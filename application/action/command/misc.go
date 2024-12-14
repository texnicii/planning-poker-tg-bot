package cmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type UnknownCommandHandler struct {
	handler.Model
	IsCallback bool
}

func (cmd UnknownCommandHandler) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	response := new(messaging.ResponseBag)
	if cmd.IsCallback {
		// TODO i18n
		response.AddCallbackResponse(update.CallbackQuery.ID, "Hm... something goes wrong 🙄")
	} else {
		response.AddChatResponse(update.Message.Chat.ID, "👀Unknown command")
	}

	return response
}

type Echo struct {
	*handler.Model
}

func (cmd Echo) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	response := new(messaging.ResponseBag)
	response.AddChatResponse(update.Message.Chat.ID, "✅")

	return response
}
