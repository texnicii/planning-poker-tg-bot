package cmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type UnknownCommandHandler struct {
	IsCallback bool
}

func (cmd UnknownCommandHandler) Handle(update tgbotapi.Update) messaging.ResponseBag {
	response := messaging.ResponseBag{}
	if cmd.IsCallback {
		// TODO i18n
		response.AddCallbackResponse(update.CallbackQuery.ID, "Hm... something goes wrong 🙄")
	} else {
		response.AddChatResponse(update.Message.Chat.ID, "👀Unknown command")
	}

	return response
}

type Echo struct {
}

func (cmd Echo) Handle(update tgbotapi.Update) messaging.ResponseBag {
	response := messaging.ResponseBag{}
	response.AddChatResponse(update.Message.Chat.ID, "✅")

	return response
}

func MakeErrorResponse(chatId int64) messaging.ResponseBag {
	response := messaging.ResponseBag{}
	response.AddChatResponse(chatId, "Hm... something goes wrong 🙄")

	return response
}
